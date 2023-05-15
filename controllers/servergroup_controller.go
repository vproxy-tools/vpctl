/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	c "github.com/vproxy-tools/vpctl/pkg/vproxy_config"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strconv"
	"sync/atomic"
	"time"

	m "github.com/vproxy-tools/vpctl/api/v1alpha1"
)

// ServerGroupReconciler reconciles a ServerGroup object
type ServerGroupReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	recorder record.EventRecorder
	cache    *ObjectUsageCache

	hcStarted atomic.Bool
	hcChannel chan *c.HealthCheckEventChannelMessage
}

//+kubebuilder:rbac:groups=app.vproxy.io,resources=servergroups,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.vproxy.io,resources=servergroups/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.vproxy.io,resources=servergroups/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ServerGroup object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ServerGroupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	logger := log.FromContext(ctx)
	o := m.ServerGroup{}
	err = r.Get(ctx, req.NamespacedName, &o)
	if err != nil {
		if errors.IsNotFound(err) {
			err = nil
			return
		}
		return
	}
	res, err = r.reconcile(ctx, logger, &o)
	if err != nil {
		r.recorder.Eventf(&o, v1.EventTypeWarning, "ReconcileFailed", "reconcile failed: %v", err)
		return
	}
	if !res.Requeue {
		r.recorder.Eventf(&o, v1.EventTypeNormal, "ReconcileSucceeded", "reconcile succeeded")
	}
	return
}

func (r *ServerGroupReconciler) reconcile(ctx context.Context, logger logr.Logger, o *m.ServerGroup) (res ctrl.Result, err error) {
	var todo []*c.Todo

	if o.DeletionTimestamp != nil {
		// need to be deleted
		todo, err = c.DeleteByConfig([]c.Config{
			&c.ServerGroup{Base: formatResourceBase(o.ObjectMeta)},
		})
	} else {
		err = addFinalizer(ctx, r.Client, o, &logger)
		if err != nil {
			return
		}
		apply := &c.ServerGroup{
			Base: formatResourceBase(o.ObjectMeta),
		}
		o.Spec.ServerGroupSelfSpec.DeepCopyInto(&apply.Spec.ServerGroupSelfSpec)
		var ss []c.StaticServer
		var skip bool
		if ss, skip, err = r.buildServers(ctx, o); err != nil || skip {
			return
		}
		apply.Spec.Servers.Static = ss
		todo, err = c.ApplyByConfig([]c.Config{apply})
	}

	if err != nil {
		return
	}
	err = c.RunTodoList(todo, &logger)
	if err != nil {
		return
	}

	if o.DeletionTimestamp != nil {
		r.cache.Remove(o)
		err = removeFinalizer(ctx, r.Client, o, &logger)
		if err != nil {
			return
		}
	} else {
		epNames := sets.Set[string]{}
		for _, ep := range o.Spec.Servers.Endpoints {
			epNames.Insert(ep.Name)
		}
		r.cache.Update(o.Namespace, epNames, o)

		r.ensureHCLoop()

		var sg *c.ServerGroup
		name := formatResourceName(o.Namespace, o.Name)
		sg, err = c.GetServerGroup(name)
		if err != nil {
			return
		}
		if sg == nil {
			logger.Error(nil, "should not happen: ServerGroup "+name+" does not exist after config applied")
		} else {
			sg.Status.DeepCopyInto(&o.Status.ServerGroupStatus)
			err = r.Client.Status().Update(ctx, o)
			if err != nil {
				logger.Error(err, "update status failed for ServerGroup "+o.Namespace+"/"+o.Name)
				return
			}
		}
	}

	return
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServerGroupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("vpctl-controller")
	r.cache = endpointsUsageCache
	r.hcChannel = make(chan *c.HealthCheckEventChannelMessage)
	return ctrl.NewControllerManagedBy(mgr).
		For(&m.ServerGroup{}).
		Complete(r)
}

func (r *ServerGroupReconciler) buildServers(ctx context.Context, o *m.ServerGroup) (ss []c.StaticServer, skip bool, err error) {
	type nameaddr struct {
		ep   string
		name string
		addr string
	}
	type tmp struct {
		ep     *v1.Endpoints
		port   int
		weight int
		size   int
		addrs  []nameaddr

		multiplier int
	}
	ls := []tmp{}
	for _, x := range o.Spec.Servers.Endpoints {
		ep := &v1.Endpoints{}
		err = r.Get(ctx, types.NamespacedName{
			Namespace: o.Namespace,
			Name:      x.Name,
		}, ep)
		if err != nil {
			if errors.IsNotFound(err) {
				r.recorder.Eventf(o, v1.EventTypeWarning, "MissingDependentResource", "missing endpoints %s/%s", o.Namespace, x.Name)
				err = nil
				skip = true
				return
			}
			return
		}
		nameaddrs := sets.Set[nameaddr]{}
		for _, s := range ep.Subsets {
			for _, a := range s.Addresses {
				if a.IP != "" {
					name := ""
					if a.TargetRef != nil {
						name = a.TargetRef.Name
					}
					nameaddrs.Insert(nameaddr{
						ep:   ep.Name,
						name: name,
						addr: a.IP,
					})
				}
			}
		}
		if nameaddrs.Len() == 0 {
			continue
		}
		var w int
		if x.Weight == nil {
			w = 1
		} else {
			w = *x.Weight
		}
		ls = append(ls, tmp{
			ep:     ep,
			port:   x.Port,
			weight: w,
			size:   len(nameaddrs),
			addrs:  nameaddrs.UnsortedList(),

			multiplier: 1,
		})
	}

	ss = []c.StaticServer{}

	if len(ls) == 0 {
		return
	}
	if len(ls) == 1 {
		w := 1
		for _, a := range ls[0].addrs {
			l4addr := a.addr + ":" + strconv.Itoa(ls[0].port)
			name := l4addr
			if a.name != "" {
				name = a.name + ":" + name
			}
			name = a.ep + ":" + name
			ss = append(ss, c.StaticServer{
				Name:    name,
				Address: l4addr,
				Weight:  &w,
			})
		}
		return
	}

	for idx, x := range ls {
		for _, y := range ls {
			if x.ep == y.ep {
				continue
			}
			x.multiplier *= y.size
		}
		ls[idx] = x
	}

	for _, e := range ls {
		for _, a := range e.addrs {
			l4addr := a.addr + ":" + strconv.Itoa(e.port)
			name := l4addr
			if a.name != "" {
				name = a.name + ":" + name
			}
			name = a.ep + ":" + name
			func(w int) {
				ss = append(ss, c.StaticServer{
					Name:    name,
					Address: l4addr,
					Weight:  &w,
				})
			}(e.weight * e.multiplier)
		}
	}

	g := -1
	for _, s := range ss {
		if *s.Weight == 0 {
			continue
		}
		if g == -1 {
			g = *s.Weight
		} else {
			g = gcd(g, *s.Weight)
		}
	}
	if g != -1 {
		for idx, s := range ss {
			func(w int) {
				ss[idx].Weight = &w
			}(*s.Weight / g)
		}
	}
	return
}

func (r *ServerGroupReconciler) ensureHCLoop() {
	if !r.hcStarted.CompareAndSwap(false, true) {
		return // already started
	}
	ctx := context.TODO()
	logger := log.FromContext(context.WithValue(ctx, "loop", "hc"))
	logger.Info("trying to start hc watching loop ...")
	stop := make(chan bool)
	err := c.WatchHealthCheck(r.hcChannel, stop)
	if err != nil {
		logger.Error(err, "failed to watch health check")
		stop <- true
		close(stop)
		r.hcStarted.Store(false)
		return
	}
	go func() {
		logger.Info("begin to watch hc events")
		defer func() { // ensure restart
			stop <- true
			close(stop)
			r.hcStarted.Store(false)

			go func() {
				time.Sleep(5 * time.Second)
				r.ensureHCLoop()
			}()
		}()

		for {
			e := <-r.hcChannel
			if e == nil || e.Err != nil {
				if e == nil {
					logger.Error(nil, "hc channel produced nil")
				} else {
					logger.Error(e.Err, "received error message from hc channel")
				}
				break
			}
			evt := e.Evt
			name := evt.ServerGroup.Metadata.Name
			if name == "" {
				logger.Error(nil, "event from hc channel does not contain server group name: "+fmt.Sprintf("%+v", evt))
				continue
			}
			ns, n, ok := extractNsName(name)
			if !ok {
				logger.Info("event from hc channel is not part of k8s: " + name)
				continue
			}
			logger.Info("event from hc channel: ServerGroup " + ns + "/" + n)
			sg := m.ServerGroup{}
			err = r.Get(ctx, types.NamespacedName{
				Namespace: ns,
				Name:      n,
			}, &sg)
			if err != nil {
				logger.Error(err, "failed to retrieve server group from k8s: "+ns+"/"+n)
				continue
			}
			sg.SyncId += 1
			err = r.Update(ctx, &sg)
			if err != nil {
				logger.Error(err, "failed to update syncId of server group: "+ns+"/"+n)
				continue
			}
		}
	}()
}

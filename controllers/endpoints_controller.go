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
	"github.com/go-logr/logr"
	m "github.com/vproxy-tools/vpctl/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// EndpointsReconciler reconciles a Endpoints object
type EndpointsReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	recorder record.EventRecorder
	cache    *ObjectUsageCache
}

//+kubebuilder:rbac:groups=vproxy.io,resources=endpoints,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=vproxy.io,resources=endpoints/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=vproxy.io,resources=endpoints/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Endpoints object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *EndpointsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	logger := log.FromContext(ctx)
	o := v1.Endpoints{}
	err = r.Get(ctx, req.NamespacedName, &o)
	if err != nil {
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

func (r *EndpointsReconciler) reconcile(ctx context.Context, logger logr.Logger, o *v1.Endpoints) (res ctrl.Result, err error) {
	if o.DeletionTimestamp != nil {
		// need to be deleted
		set := r.cache.Get(o.Namespace, o.Name)
		if set != nil && set.Len() > 0 {
			r.recorder.Eventf(o, v1.EventTypeWarning, "InUse", "the endpoints is still in use: "+setToString(&set))
			return // no error
		}
		err = removeFinalizer(ctx, r.Client, o, &logger)
		if err != nil {
			return
		}
		return
	} else {
		err = addFinalizer(ctx, r.Client, o, &logger)
		if err != nil {
			return
		}
		set := r.cache.Get(o.Namespace, o.Name)
		for v := range set {
			if v.Kind == "ServerGroup" {
				sg := m.ServerGroup{}
				err = r.Client.Get(ctx, types.NamespacedName{
					Namespace: v.Namespace,
					Name:      v.Name,
				}, &sg)
				if err != nil {
					return
				}
				sg.SyncId += 1
				err = r.Client.Update(ctx, &sg)
				if err != nil {
					return
				}
			} else {
				logger.Error(nil, "unsupported kind when triggering sync by endpoints update: "+v.String())
			}
		}
		return
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *EndpointsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.cache = endpointsUsageCache
	r.recorder = mgr.GetEventRecorderFor("vpctl-controller")
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Endpoints{}).
		Complete(r)
}

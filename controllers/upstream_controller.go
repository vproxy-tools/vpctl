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
	c "github.com/vproxy-tools/vpctl/pkg/vproxy_config"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	m "github.com/vproxy-tools/vpctl/api/v1alpha1"
)

// UpstreamReconciler reconciles a Upstream object
type UpstreamReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=vproxy.io.vproxy.io,resources=upstreams,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=vproxy.io.vproxy.io,resources=upstreams/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=vproxy.io.vproxy.io,resources=upstreams/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Upstream object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *UpstreamReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	logger := log.FromContext(ctx)
	o := m.Upstream{}
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

func (r *UpstreamReconciler) reconcile(ctx context.Context, logger logr.Logger, o *m.Upstream) (res ctrl.Result, err error) {
	var todo []*c.Todo

	if o.DeletionTimestamp != nil {
		// need to be deleted
		todo, err = c.DeleteByConfig([]c.Config{
			&c.Upstream{Base: formatResourceBase(o.ObjectMeta)},
		})
	} else {
		err = addFinalizer(ctx, r.Client, o, &logger)
		if err != nil {
			return
		}
		apply := &c.Upstream{
			Base: formatResourceBase(o.ObjectMeta),
		}
		o.Spec.UpstreamSpec.DeepCopyInto(&apply.Spec)
		for idx, sg := range apply.Spec.ServerGroups {
			addNsToResourceName(o.Namespace, &sg.Name)
			apply.Spec.ServerGroups[idx] = sg
		}
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
		err = removeFinalizer(ctx, r.Client, o, &logger)
		if err != nil {
			return
		}
	}

	return
}

// SetupWithManager sets up the controller with the Manager.
func (r *UpstreamReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("vpctl-controller")
	return ctrl.NewControllerManagedBy(mgr).
		For(&m.Upstream{}).
		Complete(r)
}

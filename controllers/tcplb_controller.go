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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/log"

	m "github.com/vproxy-tools/vpctl/api/v1alpha1"
	c "github.com/vproxy-tools/vpctl/pkg/vproxy_config"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TcpLbReconciler reconciles a TcpLb object
type TcpLbReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=app.vproxy.io,resources=tcplbs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.vproxy.io,resources=tcplbs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.vproxy.io,resources=tcplbs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TcpLb object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *TcpLbReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	logger := log.FromContext(ctx)
	o := m.TcpLb{}
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

func (r *TcpLbReconciler) reconcile(ctx context.Context, logger logr.Logger, o *m.TcpLb) (res ctrl.Result, err error) {
	var todo []*c.Todo

	if o.DeletionTimestamp != nil {
		// need to be deleted
		todo, err = c.DeleteOne("CertKey", formatResourceName(o.Namespace, o.Name))
	} else {
		err = addFinalizer(ctx, r.Client, o, &logger)
		if err != nil {
			return
		}
		apply := &c.TcpLb{
			Base: formatResourceBase(o.ObjectMeta),
		}
		o.Spec.TcpLbSpec.DeepCopyInto(&apply.Spec)
		addNsToResourceName(o.Namespace, &apply.Spec.Backend)
		addNsToResourceNameList(o.Namespace, apply.Spec.ListOfCertKey)
		addNsToResourceName(o.Namespace, &apply.Spec.SecurityGroup)
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
func (r *TcpLbReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("vpctl-controller")
	return ctrl.NewControllerManagedBy(mgr).
		For(&m.TcpLb{}).
		Complete(r)
}

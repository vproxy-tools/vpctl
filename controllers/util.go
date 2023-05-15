package controllers

import (
	"context"
	"github.com/go-logr/logr"
	c "github.com/vproxy-tools/vpctl/pkg/vproxy_config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func formatResourceName(ns, n string) string {
	return ns + "@" + n
}

func addNsToResourceName(ns string, n *string) {
	*n = ns + "@" + *n
}

func addNsToResourceNameList(ns string, nn []string) {
	for idx, n := range nn {
		nn[idx] = ns + "@" + n
	}
}

func formatResourceBase(o metav1.ObjectMeta) c.Base {
	return c.Base{
		Metadata: c.Metadata{
			Name:        formatResourceName(o.Namespace, o.Name),
			Annotations: o.Annotations,
		},
	}
}

const finalizer = "vproxy.io/vpctl"

func addFinalizer(ctx context.Context, client client.Client, o client.Object, logger *logr.Logger) error {
	finalizers := o.GetFinalizers()
	for _, f := range finalizers {
		if f == finalizer {
			return nil
		}
	}
	finalizers = append(finalizers, finalizer)
	o.SetFinalizers(finalizers)
	err := client.Update(ctx, o)
	if err != nil {
		logger.Error(err, "failed to add finalizer "+finalizer+" for "+o.GetNamespace()+"/"+o.GetName())
		return err
	}
	err = client.Get(ctx, types.NamespacedName{
		Namespace: o.GetNamespace(),
		Name:      o.GetName(),
	}, o)
	if err != nil {
		logger.Error(err, "failed to retrieve "+o.GetNamespace()+"/"+o.GetName()+" after adding finalizer")
		return err
	}
	logger.Info("finalizer " + finalizer + " added for " + o.GetNamespace() + "/" + o.GetName())
	return nil
}

func removeFinalizer(ctx context.Context, client client.Client, o client.Object, logger *logr.Logger) error {
	finalizers := o.GetFinalizers()
	if len(finalizers) == 0 {
		return nil
	}
	m := map[string]bool{}
	for _, f := range finalizers {
		m[f] = true
	}
	if !m[finalizer] {
		return nil
	}
	delete(m, finalizer)
	now := make([]string, len(m))
	idx := 0
	for k := range m {
		now[idx] = k
		idx += 1
	}
	o.SetFinalizers(now)
	err := client.Update(ctx, o)
	if err != nil {
		logger.Error(err, "failed to remove finalizer "+finalizer)
		return err
	}
	return nil
}

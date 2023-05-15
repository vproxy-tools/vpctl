package controllers

import (
	"context"
	"github.com/go-logr/logr"
	c "github.com/vproxy-tools/vpctl/pkg/vproxy_config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
	"sync"
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

func extractNsName(name string) (ns, n string, ok bool) {
	if !strings.Contains(name, "@") {
		return
	}
	split := strings.Split(name, "@")
	if len(split) != 2 {
		return
	}
	ns = strings.TrimSpace(split[0])
	n = strings.TrimSpace(split[1])
	if ns == "" || n == "" {
		return
	}
	ok = true
	return
}

func formatResourceBase(o metav1.ObjectMeta) c.Base {
	return c.Base{
		Metadata: c.Metadata{
			Name:        formatResourceName(o.Namespace, o.Name),
			Annotations: o.Annotations,
		},
	}
}

type ObjectName struct {
	Kind      string
	Namespace string
	Name      string
}

func (n ObjectName) String() string {
	return n.Kind + ":" + n.Namespace + "/" + n.Name
}

func formatObjectName(o client.Object) ObjectName {
	return ObjectName{
		Kind:      o.GetObjectKind().GroupVersionKind().Kind,
		Namespace: o.GetNamespace(),
		Name:      o.GetName(),
	}
}

type ObjectUsageCache struct {
	lock sync.RWMutex
	m    map[string]*sets.Set[ObjectName]
}

func (cache *ObjectUsageCache) Has(ns string, n string) bool {
	cache.lock.RLock()
	defer cache.lock.RUnlock()
	return cache.m[formatResourceName(ns, n)] != nil
}

func (cache *ObjectUsageCache) Get(ns string, n string) sets.Set[ObjectName] {
	cache.lock.RLock()
	defer cache.lock.RUnlock()
	ls := cache.m[formatResourceName(ns, n)]
	if ls == nil {
		return nil
	}
	return ls.Clone()
}

func (cache *ObjectUsageCache) Update(ns string, names sets.Set[string], o client.Object) {
	obj := formatObjectName(o)
	cache.lock.Lock()
	defer cache.lock.Unlock()

	for s, v := range cache.m {
		v.Delete(obj)
		if v.Len() == 0 {
			delete(cache.m, s)
		}
	}
	for n := range names {
		name := formatResourceName(ns, n)
		if cache.m[name] == nil {
			cache.m[name] = &sets.Set[ObjectName]{}
		}
		cache.m[name].Insert(obj)
	}
}

func (cache *ObjectUsageCache) Remove(o client.Object) {
	obj := formatObjectName(o)
	cache.lock.Lock()
	defer cache.lock.Unlock()

	for s, v := range cache.m {
		v.Delete(obj)
		if v.Len() == 0 {
			delete(cache.m, s)
		}
	}
}

func setToString(set *sets.Set[ObjectName]) string {
	sb := strings.Builder{}
	isFirst := true
	for set.Len() > 0 {
		if isFirst {
			isFirst = false
		} else {
			sb.WriteString(", ")
		}
		v, _ := set.PopAny()
		sb.WriteString(v.String())
	}
	return sb.String()
}

var endpointsUsageCache = &ObjectUsageCache{
	m: map[string]*sets.Set[ObjectName]{},
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

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

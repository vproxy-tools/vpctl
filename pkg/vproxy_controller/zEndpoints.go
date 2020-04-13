package vproxy_controller
import (
	"encoding/json"
	"reflect"
	"time"
)

func DefineEndpointsFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateEndpoints(meta.ns, func(resources []*Endpoints) ([]*Endpoints, bool) {
			index := -1
			for i, x := range resources {
				if x.Metadata.Name == meta.name {
					index = i
					break
				}
			}
			res := Endpoints{}
			err := json.Unmarshal(str, &res)
			if err != nil {
				Log("deserialize Endpoints failed: %v", err)
				return resources, false
			}
			if index == -1 {
				resources = append(resources, &res)
			} else {
				if reflect.DeepEqual(resources[index], &res) {
					return resources, false
				}
				resources[index] = &res
			}
			return resources, true
		})
	}
}

func DeleteEndpointsFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateEndpoints(ns, func(resource []*Endpoints) ([]*Endpoints, bool) {
			for i, x := range resource {
				if x.Metadata.Name == n {
					resource = append(resource[:i], resource[i+1:]...)
					return resource, true
				}
			}
			return resource, false
		})
	}
}

func ClearEndpointsFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateEndpoints(ns, func(resource []*Endpoints) ([]*Endpoints, bool) {
				for _, x := range resource {
					x.M.Pending = time.Now().Unix()
				}
				return resource, false
			})
		}
	}
}

func (p *_pool) UpdateEndpoints(namespace string, f func([]*Endpoints) ([]*Endpoints, bool)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	res, changed := f(rp.ep)
	if changed {
		rp.ep = res
	}
	p.trimNamespace(namespace)
	p.lock.Unlock()
	if changed {
		p.trigger()
	}
}

func (p *_pool) ClearPendingEndpoints(namespace string) {
	p.lock.RLock()
	rp := p.ensureAndGetNamespace(namespace)
	now := time.Now().Unix()
	needUpdate := false
	for _, ep := range rp.ep {
		if ep.M.Pending != 0 && now-ep.M.Pending > ClearPendingTimeoutSeconds {
			needUpdate = true
			break
		}
	}
	p.lock.RUnlock()
	if needUpdate {
		p.UpdateEndpoints(namespace, func(resource []*Endpoints) ([]*Endpoints, bool) {
			newList := make([]*Endpoints, 0)
			for _, ep := range rp.ep {
				if ep.M.Pending != 0 && now-ep.M.Pending > ClearPendingTimeoutSeconds {
					continue
				}
				newList = append(newList, ep)
			}
			return newList, true
		})
	}
}

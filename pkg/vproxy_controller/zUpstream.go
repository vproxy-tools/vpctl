package vproxy_controller
import (
	"../vproxy_config"
	"encoding/json"
	"reflect"
	"time"
)

func DefineUpstreamFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateUpstream(meta.ns, func(resources []*Upstream) ([]*Upstream, bool) {
			index := -1
			for i, x := range resources {
				if x.Metadata.Name == meta.name {
					index = i
					break
				}
			}
			res := Upstream{}
			err := json.Unmarshal(str, &res)
			if err != nil {
				Log("deserialize Upstream failed: %v", err)
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

func DeleteUpstreamFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateUpstream(ns, func(resource []*Upstream) ([]*Upstream, bool) {
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

func ClearUpstreamFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateUpstream(ns, func(resource []*Upstream) ([]*Upstream, bool) {
				for _, x := range resource {
					x.M.Pending = time.Now().Unix()
				}
				return resource, false
			})
		}
	}
}

func (p *_pool) UpdateUpstream(namespace string, f func([]*Upstream) ([]*Upstream, bool)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	res, changed := f(rp.ups)
	if changed {
		rp.ups = res
	}
	p.trimNamespace(namespace)
	p.lock.Unlock()
	if changed {
		p.trigger()
	}
}

func (p *_pool) ClearPendingUpstream(namespace string) {
	p.lock.RLock()
	rp := p.ensureAndGetNamespace(namespace)
	now := time.Now().Unix()
	needUpdate := false
	for _, ups := range rp.ups {
		if ups.M.Pending != 0 && now-ups.M.Pending > ClearPendingTimeoutSeconds {
			needUpdate = true
			break
		}
	}
	p.lock.RUnlock()
	if needUpdate {
		p.UpdateUpstream(namespace, func(resource []*Upstream) ([]*Upstream, bool) {
			newList := make([]*Upstream, 0)
			for _, ups := range rp.ups {
				if ups.M.Pending != 0 && now-ups.M.Pending > ClearPendingTimeoutSeconds {
					continue
				}
				newList = append(newList, ups)
			}
			return newList, true
		})
	}
}

func GcUpstream(configs []vproxy_config.Config) ([]*vproxy_config.Todo, error) {
	list, err := vproxy_config.ListUpstream()
	if err != nil {
		return nil, err
	}
	namesInConfigs := map[string]bool{}
	for _, c := range configs {
		namesInConfigs[c.GetBase().Metadata.Name] = true
	}
	namesToDelete := make([]string, 0)
	for _, x := range list {
		if !namesInConfigs[x.Metadata.Name] {
			namesToDelete = append(namesToDelete, x.Metadata.Name)
		}
	}
	ret := make([]*vproxy_config.Todo, 0)
	for _, name := range namesToDelete {
		todo, err := vproxy_config.DeleteOne("Upstream", name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, todo...)
	}
	return ret, nil
}

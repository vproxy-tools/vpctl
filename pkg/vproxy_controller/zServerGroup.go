package vproxy_controller
import (
	"../vproxy_config"
	"encoding/json"
	"reflect"
	"time"
)

func DefineServerGroupFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateServerGroup(meta.ns, func(resources []*ServerGroup) ([]*ServerGroup, bool) {
			index := -1
			for i, x := range resources {
				if x.Metadata.Name == meta.name {
					index = i
					break
				}
			}
			res := ServerGroup{}
			err := json.Unmarshal(str, &res)
			if err != nil {
				Log("deserialize ServerGroup failed: %v", err)
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

func DeleteServerGroupFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateServerGroup(ns, func(resource []*ServerGroup) ([]*ServerGroup, bool) {
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

func ClearServerGroupFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateServerGroup(ns, func(resource []*ServerGroup) ([]*ServerGroup, bool) {
				for _, x := range resource {
					x.M.Pending = time.Now().Unix()
				}
				return resource, false
			})
		}
	}
}

func (p *_pool) UpdateServerGroup(namespace string, f func([]*ServerGroup) ([]*ServerGroup, bool)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	res, changed := f(rp.sg)
	if changed {
		rp.sg = res
	}
	p.trimNamespace(namespace)
	p.lock.Unlock()
	if changed {
		p.trigger()
	}
}

func (p *_pool) ClearPendingServerGroup(namespace string) {
	p.lock.RLock()
	rp := p.ensureAndGetNamespace(namespace)
	now := time.Now().Unix()
	needUpdate := false
	for _, sg := range rp.sg {
		if sg.M.Pending != 0 && now-sg.M.Pending > ClearPendingTimeoutSeconds {
			needUpdate = true
			break
		}
	}
	p.lock.RUnlock()
	if needUpdate {
		p.UpdateServerGroup(namespace, func(resource []*ServerGroup) ([]*ServerGroup, bool) {
			newList := make([]*ServerGroup, 0)
			for _, sg := range rp.sg {
				if sg.M.Pending != 0 && now-sg.M.Pending > ClearPendingTimeoutSeconds {
					continue
				}
				newList = append(newList, sg)
			}
			return newList, true
		})
	}
}

func GcServerGroup(configs []vproxy_config.Config) ([]*vproxy_config.Todo, error) {
	list, err := vproxy_config.ListServerGroup()
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
		todo, err := vproxy_config.DeleteOne("ServerGroup", name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, todo...)
	}
	return ret, nil
}

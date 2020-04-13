package vproxy_controller
import (
	"../vproxy_config"
	"encoding/json"
	"reflect"
	"time"
)

func DefineSocks5ServerFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateSocks5Server(meta.ns, func(resources []*Socks5Server) ([]*Socks5Server, bool) {
			index := -1
			for i, x := range resources {
				if x.Metadata.Name == meta.name {
					index = i
					break
				}
			}
			res := Socks5Server{}
			err := json.Unmarshal(str, &res)
			if err != nil {
				Log("deserialize Socks5Server failed: %v", err)
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

func DeleteSocks5ServerFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateSocks5Server(ns, func(resource []*Socks5Server) ([]*Socks5Server, bool) {
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

func ClearSocks5ServerFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateSocks5Server(ns, func(resource []*Socks5Server) ([]*Socks5Server, bool) {
				for _, x := range resource {
					x.M.Pending = time.Now().Unix()
				}
				return resource, false
			})
		}
	}
}

func (p *_pool) UpdateSocks5Server(namespace string, f func([]*Socks5Server) ([]*Socks5Server, bool)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	res, changed := f(rp.socks5)
	if changed {
		rp.socks5 = res
	}
	p.trimNamespace(namespace)
	p.lock.Unlock()
	if changed {
		p.trigger()
	}
}

func (p *_pool) ClearPendingSocks5Server(namespace string) {
	p.lock.RLock()
	rp := p.ensureAndGetNamespace(namespace)
	now := time.Now().Unix()
	needUpdate := false
	for _, socks5 := range rp.socks5 {
		if socks5.M.Pending != 0 && now-socks5.M.Pending > ClearPendingTimeoutSeconds {
			needUpdate = true
			break
		}
	}
	p.lock.RUnlock()
	if needUpdate {
		p.UpdateSocks5Server(namespace, func(resource []*Socks5Server) ([]*Socks5Server, bool) {
			newList := make([]*Socks5Server, 0)
			for _, socks5 := range rp.socks5 {
				if socks5.M.Pending != 0 && now-socks5.M.Pending > ClearPendingTimeoutSeconds {
					continue
				}
				newList = append(newList, socks5)
			}
			return newList, true
		})
	}
}

func GcSocks5Server(configs []vproxy_config.Config) ([]*vproxy_config.Todo, error) {
	list, err := vproxy_config.ListSocks5Server()
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
		todo, err := vproxy_config.DeleteOne("Socks5Server", name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, todo...)
	}
	return ret, nil
}

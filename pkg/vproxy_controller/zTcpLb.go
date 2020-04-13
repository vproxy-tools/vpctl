package vproxy_controller
import (
	"../vproxy_config"
	"encoding/json"
	"reflect"
	"time"
)

func DefineTcpLbFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateTcpLb(meta.ns, func(resources []*TcpLb) ([]*TcpLb, bool) {
			index := -1
			for i, x := range resources {
				if x.Metadata.Name == meta.name {
					index = i
					break
				}
			}
			res := TcpLb{}
			err := json.Unmarshal(str, &res)
			if err != nil {
				Log("deserialize TcpLb failed: %v", err)
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

func DeleteTcpLbFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateTcpLb(ns, func(resource []*TcpLb) ([]*TcpLb, bool) {
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

func ClearTcpLbFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateTcpLb(ns, func(resource []*TcpLb) ([]*TcpLb, bool) {
				for _, x := range resource {
					x.M.Pending = time.Now().Unix()
				}
				return resource, false
			})
		}
	}
}

func (p *_pool) UpdateTcpLb(namespace string, f func([]*TcpLb) ([]*TcpLb, bool)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	res, changed := f(rp.tl)
	if changed {
		rp.tl = res
	}
	p.trimNamespace(namespace)
	p.lock.Unlock()
	if changed {
		p.trigger()
	}
}

func (p *_pool) ClearPendingTcpLb(namespace string) {
	p.lock.RLock()
	rp := p.ensureAndGetNamespace(namespace)
	now := time.Now().Unix()
	needUpdate := false
	for _, tl := range rp.tl {
		if tl.M.Pending != 0 && now-tl.M.Pending > ClearPendingTimeoutSeconds {
			needUpdate = true
			break
		}
	}
	p.lock.RUnlock()
	if needUpdate {
		p.UpdateTcpLb(namespace, func(resource []*TcpLb) ([]*TcpLb, bool) {
			newList := make([]*TcpLb, 0)
			for _, tl := range rp.tl {
				if tl.M.Pending != 0 && now-tl.M.Pending > ClearPendingTimeoutSeconds {
					continue
				}
				newList = append(newList, tl)
			}
			return newList, true
		})
	}
}

func GcTcpLb(configs []vproxy_config.Config) ([]*vproxy_config.Todo, error) {
	list, err := vproxy_config.ListTcpLb()
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
		todo, err := vproxy_config.DeleteOne("TcpLb", name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, todo...)
	}
	return ret, nil
}

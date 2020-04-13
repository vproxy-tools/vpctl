package vproxy_controller
import (
	"../vproxy_config"
	"encoding/json"
	"reflect"
	"time"
)

func DefineSecurityGroupFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateSecurityGroup(meta.ns, func(resources []*SecurityGroup) ([]*SecurityGroup, bool) {
			index := -1
			for i, x := range resources {
				if x.Metadata.Name == meta.name {
					index = i
					break
				}
			}
			res := SecurityGroup{}
			err := json.Unmarshal(str, &res)
			if err != nil {
				Log("deserialize SecurityGroup failed: %v", err)
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

func DeleteSecurityGroupFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateSecurityGroup(ns, func(resource []*SecurityGroup) ([]*SecurityGroup, bool) {
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

func ClearSecurityGroupFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateSecurityGroup(ns, func(resource []*SecurityGroup) ([]*SecurityGroup, bool) {
				for _, x := range resource {
					x.M.Pending = time.Now().Unix()
				}
				return resource, false
			})
		}
	}
}

func (p *_pool) UpdateSecurityGroup(namespace string, f func([]*SecurityGroup) ([]*SecurityGroup, bool)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	res, changed := f(rp.secg)
	if changed {
		rp.secg = res
	}
	p.trimNamespace(namespace)
	p.lock.Unlock()
	if changed {
		p.trigger()
	}
}

func (p *_pool) ClearPendingSecurityGroup(namespace string) {
	p.lock.RLock()
	rp := p.ensureAndGetNamespace(namespace)
	now := time.Now().Unix()
	needUpdate := false
	for _, secg := range rp.secg {
		if secg.M.Pending != 0 && now-secg.M.Pending > ClearPendingTimeoutSeconds {
			needUpdate = true
			break
		}
	}
	p.lock.RUnlock()
	if needUpdate {
		p.UpdateSecurityGroup(namespace, func(resource []*SecurityGroup) ([]*SecurityGroup, bool) {
			newList := make([]*SecurityGroup, 0)
			for _, secg := range rp.secg {
				if secg.M.Pending != 0 && now-secg.M.Pending > ClearPendingTimeoutSeconds {
					continue
				}
				newList = append(newList, secg)
			}
			return newList, true
		})
	}
}

func GcSecurityGroup(configs []vproxy_config.Config) ([]*vproxy_config.Todo, error) {
	list, err := vproxy_config.ListSecurityGroup()
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
		todo, err := vproxy_config.DeleteOne("SecurityGroup", name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, todo...)
	}
	return ret, nil
}

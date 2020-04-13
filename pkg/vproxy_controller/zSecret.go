package vproxy_controller
import (
	"../vproxy_config"
	"encoding/json"
	"reflect"
	"time"
)

func DefineSecretFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateSecret(meta.ns, func(resources []*Secret) ([]*Secret, bool) {
			index := -1
			for i, x := range resources {
				if x.Metadata.Name == meta.name {
					index = i
					break
				}
			}
			res := Secret{}
			err := json.Unmarshal(str, &res)
			if err != nil {
				Log("deserialize Secret failed: %v", err)
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

func DeleteSecretFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateSecret(ns, func(resource []*Secret) ([]*Secret, bool) {
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

func ClearSecretFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateSecret(ns, func(resource []*Secret) ([]*Secret, bool) {
				for _, x := range resource {
					x.M.Pending = time.Now().Unix()
				}
				return resource, false
			})
		}
	}
}

func (p *_pool) UpdateSecret(namespace string, f func([]*Secret) ([]*Secret, bool)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	res, changed := f(rp.ck)
	if changed {
		rp.ck = res
	}
	p.trimNamespace(namespace)
	p.lock.Unlock()
	if changed {
		p.trigger()
	}
}

func (p *_pool) ClearPendingSecret(namespace string) {
	p.lock.RLock()
	rp := p.ensureAndGetNamespace(namespace)
	now := time.Now().Unix()
	needUpdate := false
	for _, ck := range rp.ck {
		if ck.M.Pending != 0 && now-ck.M.Pending > ClearPendingTimeoutSeconds {
			needUpdate = true
			break
		}
	}
	p.lock.RUnlock()
	if needUpdate {
		p.UpdateSecret(namespace, func(resource []*Secret) ([]*Secret, bool) {
			newList := make([]*Secret, 0)
			for _, ck := range rp.ck {
				if ck.M.Pending != 0 && now-ck.M.Pending > ClearPendingTimeoutSeconds {
					continue
				}
				newList = append(newList, ck)
			}
			return newList, true
		})
	}
}

func GcCertKey(configs []vproxy_config.Config) ([]*vproxy_config.Todo, error) {
	list, err := vproxy_config.ListCertKey()
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
		todo, err := vproxy_config.DeleteOne("CertKey", name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, todo...)
	}
	return ret, nil
}

package vproxy_controller
import (
	"../vproxy_config"
	"encoding/json"
)
func DefineServerGroupFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateServerGroup(meta.ns, func(resources []*ServerGroup) ([]*ServerGroup, int) {
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
				return resources, -1
			}
			res.M.Version = meta.ver
			if index == -1 {
				resources = append(resources, &res)
			} else {
				resources[index] = &res
			}
			return resources, meta.ver
		})
	}
}

func DeleteServerGroupFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateServerGroup(ns, func(resource []*ServerGroup) ([]*ServerGroup, int) {
			for i, x := range resource {
				if x.Metadata.Name == n {
					resource = append(resource[:i], resource[i+1:]...)
					break
				}
			}
			return resource, -1
		})
	}
}

func ClearServerGroupFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateServerGroup(ns, func(resource []*ServerGroup) ([]*ServerGroup, int) {
				for _, x := range resource {
					x.M.Pending = true
				}
				return resource, -1
			})
		}
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

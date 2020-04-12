package vproxy_controller
import (
	"../vproxy_config"
	"encoding/json"
)
func DefineSecurityGroupFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateSecurityGroup(meta.ns, func(resources []*SecurityGroup) ([]*SecurityGroup, int) {
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

func DeleteSecurityGroupFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateSecurityGroup(ns, func(resource []*SecurityGroup) ([]*SecurityGroup, int) {
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

func ClearSecurityGroupFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateSecurityGroup(ns, func(resource []*SecurityGroup) ([]*SecurityGroup, int) {
				for _, x := range resource {
					x.M.Pending = true
				}
				return resource, -1
			})
		}
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

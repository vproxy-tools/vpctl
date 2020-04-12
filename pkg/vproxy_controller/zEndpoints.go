package vproxy_controller
import (
	"encoding/json"
)
func DefineEndpointsFunc(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.UpdateEndpoints(meta.ns, func(resources []*Endpoints) ([]*Endpoints, int) {
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

func DeleteEndpointsFunc(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.UpdateEndpoints(ns, func(resource []*Endpoints) ([]*Endpoints, int) {
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

func ClearEndpointsFunc(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.UpdateEndpoints(ns, func(resource []*Endpoints) ([]*Endpoints, int) {
				for _, x := range resource {
					x.M.Pending = true
				}
				return resource, -1
			})
		}
	}
}

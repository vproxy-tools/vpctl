#!/usr/bin/env python

templatePackage="""
package vproxy_controller
"""

templateImport1="""
import (
	"../vproxy_config"
	"encoding/json"
)
"""

templateImport2="""
import (
	"encoding/json"
)
"""

template1="""
func Define{{Kind}}Func(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.Update{{Kind}}(meta.ns, func(resources []*{{Kind}}) ([]*{{Kind}}, int) {
			index := -1
			for i, x := range resources {
				if x.Metadata.Name == meta.name {
					index = i
					break
				}
			}
			res := {{Kind}}{}
			err := json.Unmarshal(str, &res)
			if err != nil {
				Log("deserialize {{Kind}} failed: %v", err)
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

func Delete{{Kind}}Func(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.Update{{Kind}}(ns, func(resource []*{{Kind}}) ([]*{{Kind}}, int) {
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

func Clear{{Kind}}Func(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.Update{{Kind}}(ns, func(resource []*{{Kind}}) ([]*{{Kind}}, int) {
				for _, x := range resource {
					x.M.Pending = true
				}
				return resource, -1
			})
		}
	}
}
"""

template2="""
func Gc{{VProxyType}}(configs []vproxy_config.Config) ([]*vproxy_config.Todo, error) {
	list, err := vproxy_config.List{{VProxyType}}()
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
		todo, err := vproxy_config.DeleteOne("{{VProxyType}}", name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, todo...)
	}
	return ret, nil
}
"""

toGenerate = [
    {
        "file": "TcpLb",
        "type": "TcpLb",
        "kind": "TcpLb",
    },
    {
        "file": "Socks5Server",
        "type": "Socks5Server",
        "kind": "Socks5Server",
    },
    {
        "file": "DnsServer",
        "type": "DnsServer",
        "kind": "DnsServer",
    },
    {
        "file": "Upstream",
        "type": "Upstream",
        "kind": "Upstream",
    },
    {
        "file": "ServerGroup",
        "type": "ServerGroup",
        "kind": "ServerGroup",
    },
    {
        "file": "SecurityGroup",
        "type": "SecurityGroup",
        "kind": "SecurityGroup",
    },
    {
        "file": "Secret",
        "type": "CertKey",
        "kind": "Secret",
    },
    {
        "file": "Endpoints",
        "kind": "Endpoints",
        "noGc": True,
    },
]

for g in toGenerate:
    f = open('./pkg/vproxy_controller/z' + g['file'] + '.go', 'w+')
    content = templatePackage.strip() + '\n'
    if not g.has_key('noGc') or not g['noGc']:
        content += templateImport1.strip() + '\n'
    else:
        content += templateImport2.strip() + '\n'
    content += template1.replace('{{Kind}}', g['kind']).strip() + '\n'
    if not g.has_key('noGc') or not g['noGc']:
        content += template2.replace('{{VProxyType}}', g['type']).replace('{{Kind}}', g['kind']).strip() + '\n'
    f.write(content)
    f.close()

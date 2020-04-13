#!/usr/bin/env python

templatePackage="""
package vproxy_controller
"""

templateImport1="""
import (
	"../vproxy_config"
	"encoding/json"
	"reflect"
	"time"
)
"""

templateImport2="""
import (
	"encoding/json"
	"reflect"
	"time"
)
"""

template1="""
func Define{{Kind}}Func(pool Pool) func([]byte, *meta) {
	return func(str []byte, meta *meta) {
		pool.Update{{Kind}}(meta.ns, func(resources []*{{Kind}}) ([]*{{Kind}}, bool) {
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

func Delete{{Kind}}Func(pool Pool) func(string, string) {
	return func(ns string, n string) {
		pool.Update{{Kind}}(ns, func(resource []*{{Kind}}) ([]*{{Kind}}, bool) {
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

func Clear{{Kind}}Func(pool Pool) func() {
	return func() {
		namespaces := pool.GetNamespaces()
		for _, ns := range namespaces {
			pool.Update{{Kind}}(ns, func(resource []*{{Kind}}) ([]*{{Kind}}, bool) {
				for _, x := range resource {
					x.M.Pending = time.Now().Unix()
				}
				return resource, false
			})
		}
	}
}

func (p *_pool) Update{{Kind}}(namespace string, f func([]*{{Kind}}) ([]*{{Kind}}, bool)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	res, changed := f(rp.{{Short}})
	if changed {
		rp.{{Short}} = res
	}
	p.trimNamespace(namespace)
	p.lock.Unlock()
	if changed {
		p.trigger()
	}
}

func (p *_pool) ClearPending{{Kind}}(namespace string) {
	p.lock.RLock()
	rp := p.ensureAndGetNamespace(namespace)
	now := time.Now().Unix()
	needUpdate := false
	for _, {{Short}} := range rp.{{Short}} {
		if {{Short}}.M.Pending != 0 && now-{{Short}}.M.Pending > ClearPendingTimeoutSeconds {
			needUpdate = true
			break
		}
	}
	p.lock.RUnlock()
	if needUpdate {
		p.Update{{Kind}}(namespace, func(resource []*{{Kind}}) ([]*{{Kind}}, bool) {
			newList := make([]*{{Kind}}, 0)
			for _, {{Short}} := range rp.{{Short}} {
				if {{Short}}.M.Pending != 0 && now-{{Short}}.M.Pending > ClearPendingTimeoutSeconds {
					continue
				}
				newList = append(newList, {{Short}})
			}
			return newList, true
		})
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
        "short": "tl",
    },
    {
        "file": "Socks5Server",
        "type": "Socks5Server",
        "kind": "Socks5Server",
        "short": "socks5",
    },
    {
        "file": "DnsServer",
        "type": "DnsServer",
        "kind": "DnsServer",
        "short": "dns",
    },
    {
        "file": "Upstream",
        "type": "Upstream",
        "kind": "Upstream",
        "short": "ups",
    },
    {
        "file": "ServerGroup",
        "type": "ServerGroup",
        "kind": "ServerGroup",
        "short": "sg",
    },
    {
        "file": "SecurityGroup",
        "type": "SecurityGroup",
        "kind": "SecurityGroup",
        "short": "secg",
    },
    {
        "file": "Secret",
        "type": "CertKey",
        "kind": "Secret",
        "short": "ck",
    },
    {
        "file": "Endpoints",
        "kind": "Endpoints",
        "noGc": True,
        "short": "ep",
    },
]

for g in toGenerate:
    f = open('./pkg/vproxy_controller/z' + g['file'] + '.go', 'w+')
    content = templatePackage.strip() + '\n'
    if not g.has_key('noGc') or not g['noGc']:
        content += templateImport1.strip() + '\n'
    else:
        content += templateImport2.strip() + '\n'
    content += '\n' + template1.replace('{{Kind}}', g['kind']).replace('{{Short}}', g['short']).strip() + '\n'
    if not g.has_key('noGc') or not g['noGc']:
        content += '\n' + template2.replace('{{VProxyType}}', g['type']).replace('{{Kind}}', g['kind']).strip() + '\n'
    f.write(content)
    f.close()

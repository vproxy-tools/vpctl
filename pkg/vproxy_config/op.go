package vproxy_config

import (
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"reflect"
	"sort"
	"strconv"
	client "vproxy_client"
	"vproxy_client/cert_key"
	"vproxy_client/dns_server"
	"vproxy_client/security_group"
	"vproxy_client/server_group"
	"vproxy_client/socks5_server"
	"vproxy_client/tcp_lb"
	"vproxy_client/upstream"
)

func getClient() *client.Vproxy {
	return client.New(httptransport.New(
		GetHost()+":"+strconv.FormatInt(int64(GetHttpPort()), 10),
		client.DefaultBasePath,
		client.DefaultSchemes,
	), strfmt.Default)
}

type Op string

const (
	OpCreate Op = "create"
	OpUpdate Op = "update"
	OpDelete Op = "delete"
	OpNone   Op = "none"
	Op404    Op = "404"
	OpText   Op = "syntax"
)

type Todo struct {
	client *client.Vproxy
	op     Op
	Ref    string
	parent interface{}
	object interface{}
	F      func(todo *Todo) error
}

func applyOneConfig(config Config, client *client.Vproxy) ([]*Todo, error) {
	if v, ok := config.(*TcpLb); ok {
		return applyTcpLb(v, client)
	} else if v, ok := config.(*Socks5Server); ok {
		return applySocks5Server(v, client)
	} else if v, ok := config.(*DnsServer); ok {
		return applyDnsServer(v, client)
	} else if v, ok := config.(*Upstream); ok {
		return applyUpstream(v, client)
	} else if v, ok := config.(*ServerGroup); ok {
		return applyServerGroup(v, client)
	} else if v, ok := config.(*SecurityGroup); ok {
		return applySecurityGroup(v, client)
	} else if v, ok := config.(*CertKey); ok {
		return applyCertKey(v, client)
	} else {
		return nil, fmt.Errorf("BUG: unknown config %s", config)
	}
}

func buildNoneOp(ref string) ([]*Todo, error) {
	return []*Todo{
		{
			client: nil,
			op:     OpNone,
			Ref:    ref,
			object: nil,
			F: func(todo *Todo) error {
				return nil // do nothing
			},
		},
	}, nil
}

func build404Op(ref string) ([]*Todo, error) {
	return []*Todo{
		{
			client: nil,
			op:     Op404,
			Ref:    ref,
			object: nil,
			F: func(todo *Todo) error {
				return nil // do nothing
			},
		},
	}, nil
}

func buildTextOp(ref string) *Todo {
	return &Todo{
		client: nil,
		op:     OpText,
		Ref:    ref,
		object: nil,
		F: func(todo *Todo) error {
			return nil // do nothing
		},
	}
}

func applyTcpLb(tl *TcpLb, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", "TcpLb", tl.Metadata.Name)
	old, err := GetTcpLb(tl.Metadata.Name)
	if err != nil {
		if _, ok := err.(*tcp_lb.GetTCPLbNotFound); ok {
			// should create
			return []*Todo{
				{
					client: client,
					op:     OpCreate,
					Ref:    ref,
					object: tl,
					F:      createTcpLb,
				},
			}, nil
		}
		return nil, err
	}
	update, err := tl.validateForUpdating(old)
	if err != nil {
		return nil, err
	}
	if update {
		return []*Todo{
			{
				client: client,
				op:     OpUpdate,
				Ref:    ref,
				object: tl,
				F:      updateTcpLb,
			},
		}, nil
	} else {
		return buildNoneOp(ref)
	}
}

func applySocks5Server(socks5 *Socks5Server, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", "Socks5Server", socks5.Metadata.Name)
	old, err := GetSocks5Server(socks5.Metadata.Name)
	if err != nil {
		if _, ok := err.(*socks5_server.GetSocks5ServerNotFound); ok {
			// should create
			return []*Todo{
				{
					client: client,
					op:     OpCreate,
					Ref:    ref,
					object: socks5,
					F:      createSocks5Server,
				},
			}, nil
		}
		return nil, err
	}
	update, err := socks5.validateForUpdating(old)
	if err != nil {
		return nil, err
	}
	if update {
		return []*Todo{
			{
				client: client,
				op:     OpUpdate,
				Ref:    ref,
				object: socks5,
				F:      updateSocks5Server,
			},
		}, nil
	} else {
		return buildNoneOp(ref)
	}
}

func applyDnsServer(dns *DnsServer, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", "DnsServer", dns.Metadata.Name)
	old, err := GetDnsServer(dns.Metadata.Name)
	if err != nil {
		if _, ok := err.(*dns_server.GetDNSServerNotFound); ok {
			// should create
			return []*Todo{
				{
					client: client,
					op:     OpCreate,
					Ref:    ref,
					object: dns,
					F:      createDnsServer,
				},
			}, nil
		}
		return nil, err
	}
	update, err := dns.validateForUpdating(old)
	if err != nil {
		return nil, err
	}
	if update {
		return []*Todo{
			{
				client: client,
				op:     OpUpdate,
				Ref:    ref,
				object: dns,
				F:      updateDnsServer,
			},
		}, nil
	} else {
		return buildNoneOp(ref)
	}
}

func getServerGroupInUpstream(ups *Upstream, sample *ServerGroupInUpstream) *ServerGroupInUpstream {
	for _, x := range ups.Spec.ServerGroups {
		if x.Name == sample.Name {
			return &x
		}
	}
	return nil
}

func buildServerGroupUpdateList(client *client.Vproxy, now *Upstream, old *Upstream, ret []*Todo) []*Todo {
	// in old
	for _, oldSg := range old.Spec.ServerGroups {
		if nil == getServerGroupInUpstream(now, &oldSg) {
			ref := fmt.Sprintf("%s:%s/%s", "Upstream/ServerGroup", old.Metadata.Name, oldSg.Name)
			x := oldSg
			// in old not in now
			// should delete
			ret = append(ret, &Todo{
				client: client,
				op:     OpDelete,
				Ref:    ref,
				parent: old,
				object: &x,
				F:      deleteServerGroupInUpstream,
			})
		}
	}
	// in now
	for _, nowSg := range now.Spec.ServerGroups {
		ref := fmt.Sprintf("%s:%s/%s", "Upstream/ServerGroup", now.Metadata.Name, nowSg.Name)
		oldSg := getServerGroupInUpstream(old, &nowSg)
		x := nowSg
		if oldSg == nil {
			// should create
			ret = append(ret, &Todo{
				client: client,
				op:     OpCreate,
				Ref:    ref,
				parent: now,
				object: &x,
				F:      createServerGroupInUpstream,
			})
		} else {
			if nowSg.validateForUpdating(oldSg) {
				// should update
				ret = append(ret, &Todo{
					client: client,
					op:     OpUpdate,
					Ref:    ref,
					parent: now,
					object: &x,
					F:      updateServerGroupInUpstream,
				})
			}
		}

	}
	return ret
}

func applyUpstream(ups *Upstream, client *client.Vproxy) ([]*Todo, error) {
	refTop := fmt.Sprintf("%s:%s", "Upstream", ups.Metadata.Name)
	old, err := GetUpstream(ups.Metadata.Name)
	if err != nil {
		if _, ok := err.(*upstream.GetUpstreamNotFound); ok {
			// should create and add all server groups
			ret := make([]*Todo, 1+len(ups.Spec.ServerGroups))
			ret[0] = &Todo{
				client: client,
				op:     OpCreate,
				Ref:    refTop,
				object: ups,
				F:      createUpstream,
			}
			for idx, sg := range ups.Spec.ServerGroups {
				ref := fmt.Sprintf("%s:%s/%s", "Upstream/ServerGroup", ups.Metadata.Name, sg.Name)
				x := sg
				ret[idx+1] = &Todo{
					client: client,
					op:     OpCreate,
					Ref:    ref,
					parent: ups,
					object: &x,
					F:      createServerGroupInUpstream,
				}
			}
			return ret, nil
		}
		return nil, err
	}

	updateTop, updateSub, err := ups.validateForUpdating(old)
	if updateTop || updateSub {
		var ret []*Todo
		if updateTop {
			ret = []*Todo{
				{
					client: client,
					op:     OpUpdate,
					Ref:    refTop,
					object: ups,
					F:      updateUpstream,
				},
			}
		} else {
			ret = make([]*Todo, 0)
		}
		if updateSub {
			ret = buildServerGroupUpdateList(client, ups, old, ret)
		}
		if len(ret) == 0 {
			ret = append(ret, buildTextOp(refTop))
		}
		return ret, nil
	} else {
		return buildNoneOp(refTop)
	}
}

func getStaticServerInServerGroup(sg *ServerGroup, sample *StaticServer) *StaticServer {
	for _, x := range sg.Spec.Servers.Static {
		if x.Name == sample.Name && x.Address == sample.Address {
			return &x
		}
	}
	return nil
}

func buildStaticServerInServerGroupUpdateList(client *client.Vproxy, now *ServerGroup, old *ServerGroup, ret []*Todo) []*Todo {
	// in old
	for _, oldSvr := range old.Spec.Servers.Static {
		if nil == getStaticServerInServerGroup(now, &oldSvr) {
			ref := fmt.Sprintf("%s:%s/%s", "ServerGroup/Server", old.Metadata.Name, oldSvr.Name)
			x := oldSvr
			// in old not in now
			// should delete
			ret = append(ret, &Todo{
				client: client,
				op:     OpDelete,
				Ref:    ref,
				parent: old,
				object: &x,
				F:      deleteStaticServerInServerGroup,
			})
		}
	}
	// in now
	for _, nowSvr := range now.Spec.Servers.Static {
		ref := fmt.Sprintf("%s:%s/%s", "ServerGroup/Server", now.Metadata.Name, nowSvr.Name)
		oldSvr := getStaticServerInServerGroup(old, &nowSvr)
		x := nowSvr
		if oldSvr == nil {
			// should create
			ret = append(ret, &Todo{
				client: client,
				op:     OpCreate,
				Ref:    ref,
				parent: now,
				object: &x,
				F:      createStaticServerInServerGroup,
			})
		} else {
			if nowSvr.validateForUpdating(oldSvr) {
				// should update
				ret = append(ret, &Todo{
					client: client,
					op:     OpUpdate,
					Ref:    ref,
					parent: now,
					object: &x,
					F:      updateStaticServerInServerGroup,
				})
			}
		}

	}
	return ret
}

func applyServerGroup(sg *ServerGroup, client *client.Vproxy) ([]*Todo, error) {
	refTop := fmt.Sprintf("%s:%s", "ServerGroup", sg.Metadata.Name)
	old, err := GetServerGroup(sg.Metadata.Name)
	if err != nil {
		if _, ok := err.(*server_group.GetServerGroupNotFound); ok {
			// should create and add all servers
			ret := make([]*Todo, 1+len(sg.Spec.Servers.Static))
			ret[0] = &Todo{
				client: client,
				op:     OpCreate,
				Ref:    refTop,
				object: sg,
				F:      createServerGroup,
			}
			for idx, svr := range sg.Spec.Servers.Static {
				refSub := fmt.Sprintf("%s:%s/%s", "ServerGroup/Server", sg.Metadata.Name, svr.Name)
				x := svr
				ret[idx+1] = &Todo{
					client: client,
					op:     OpCreate,
					Ref:    refSub,
					parent: sg,
					object: &x,
					F:      createStaticServerInServerGroup,
				}
			}
			return ret, nil
		}
		return nil, err
	}

	updateTop, updateSub, err := sg.validateForUpdating(old)
	if updateTop || updateSub {
		var ret []*Todo
		if updateTop {
			ret = []*Todo{
				{
					client: client,
					op:     OpUpdate,
					Ref:    refTop,
					object: sg,
					F:      updateServerGroup,
				},
			}
		} else {
			ret = make([]*Todo, 0)
		}
		if updateSub {
			ret = buildStaticServerInServerGroupUpdateList(client, sg, old, ret)
		}
		if len(ret) == 0 {
			ret = append(ret, buildTextOp(refTop))
		}
		return ret, nil
	} else {
		return buildNoneOp(refTop)
	}
}

func getSecurityGroupRuleInSecurityGroup(secg *SecurityGroup, sample *SecurityGroupRule) *SecurityGroupRule {
	for _, x := range secg.Spec.Rules {
		if reflect.DeepEqual(&x, sample) { // all fields immutable
			return &x
		}
	}
	return nil
}

func buildSecurityGroupRuleInSecurityGroupUpdateList(client *client.Vproxy, now *SecurityGroup, old *SecurityGroup, ret []*Todo) []*Todo {
	// in old
	for _, oldRule := range old.Spec.Rules {
		if nil == getSecurityGroupRuleInSecurityGroup(now, &oldRule) {
			ref := fmt.Sprintf("%s:%s/%s", "SecurityGroup/SecurityGroupRule", old.Metadata.Name, oldRule.Name)
			x := oldRule
			// in old not in now
			// should delete
			ret = append(ret, &Todo{
				client: client,
				op:     OpDelete,
				Ref:    ref,
				parent: old,
				object: &x,
				F:      deleteSecurityGroupRuleInSecurityGroup,
			})
		}
	}
	// in now
	for _, nowRule := range now.Spec.Rules {
		ref := fmt.Sprintf("%s:%s/%s", "SecurityGroup/SecurityGroupRule", now.Metadata.Name, nowRule.Name)
		oldRule := getSecurityGroupRuleInSecurityGroup(old, &nowRule)
		x := nowRule
		if oldRule == nil {
			// should create
			ret = append(ret, &Todo{
				client: client,
				op:     OpCreate,
				Ref:    ref,
				parent: now,
				object: &x,
				F:      createSecurityGroupRuleInSecurityGroup,
			})
		} else {
			if nowRule.validateForUpdating(oldRule) {
				// should update
				ret = append(ret, &Todo{
					client: client,
					op:     OpUpdate,
					Ref:    ref,
					parent: now,
					object: &x,
					F:      updateSecurityGroupRuleInSecurityGroup,
				})
			}
		}

	}
	return ret
}

func applySecurityGroup(secg *SecurityGroup, client *client.Vproxy) ([]*Todo, error) {
	refTop := fmt.Sprintf("%s:%s", "SecurityGroup", secg.Metadata.Name)
	old, err := GetSecurityGroup(secg.Metadata.Name)
	if err != nil {
		if _, ok := err.(*security_group.GetSecurityGroupNotFound); ok {
			// should create and add all servers
			ret := make([]*Todo, 1+len(secg.Spec.Rules))
			ret[0] = &Todo{
				client: client,
				op:     OpCreate,
				Ref:    refTop,
				object: secg,
				F:      createSecurityGroup,
			}
			for idx, secgr := range secg.Spec.Rules {
				refSub := fmt.Sprintf("%s:%s/%s", "SecurityGroup/SecurityGroupRule", secg.Metadata.Name, secgr.Name)
				x := secgr
				ret[idx+1] = &Todo{
					client: client,
					op:     OpCreate,
					Ref:    refSub,
					parent: secg,
					object: &x,
					F:      createSecurityGroupRuleInSecurityGroup,
				}
			}
			return ret, nil
		}
		return nil, err
	}

	updateTop, updateSub, err := secg.validateForUpdating(old)
	if updateTop || updateSub {
		var ret []*Todo
		if updateTop {
			ret = []*Todo{
				{
					client: client,
					op:     OpUpdate,
					Ref:    refTop,
					object: secg,
					F:      updateSecurityGroup,
				},
			}
		} else {
			ret = make([]*Todo, 0)
		}
		if updateSub {
			ret = buildSecurityGroupRuleInSecurityGroupUpdateList(client, secg, old, ret)
		}
		if len(ret) == 0 {
			ret = append(ret, buildTextOp(refTop))
		}
		return ret, nil
	} else {
		return buildNoneOp(refTop)
	}
}

func applyCertKey(ck *CertKey, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", "CertKey", ck.Metadata.Name)
	old, err := GetCertKey(ck.Metadata.Name)
	if err != nil {
		if _, ok := err.(*cert_key.DescribeCertKeyNotFound); ok {
			// should create
			return []*Todo{
				{
					client: client,
					op:     OpCreate,
					Ref:    ref,
					object: ck,
					F:      createCertKey,
				},
			}, nil
		}
		return nil, err
	}
	update, err := ck.validateForUpdating(old)
	if err != nil {
		return nil, err
	}
	if update {
		return []*Todo{
			{
				client: client,
				op:     OpUpdate,
				Ref:    ref,
				object: ck,
				F:      updateCertKey,
			},
		}, nil
	} else {
		return buildNoneOp(ref)
	}
}

func ApplyByConfig(configs []Config) ([]*Todo, error) {
	// validate
	for idx, c := range configs {
		err := c.Validate()
		if err != nil {
			return nil, fmt.Errorf("invalid document at index %d: %s", idx, err.Error())
		}
	}
	// sort
	tmp := ConfigSortingArrayForApplying(configs)
	sort.Sort(tmp)
	configs = tmp
	// check
	cli := getClient()
	todoList := make([]*Todo, 0)
	for _, c := range configs {
		t, err := applyOneConfig(c, cli)
		if err != nil {
			return nil, fmt.Errorf("cannot apply %s:%s: %s", c.GetBase().Kind, c.GetBase().Metadata.Name, err.Error())
		}
		todoList = append(todoList, t...)
	}
	return todoList, nil
}

func ApplyByFile(filename string) ([]*Todo, error) {
	configs, err := ParseFile(filename)
	if err != nil {
		return nil, err
	}
	return ApplyByConfig(configs)
}

func deleteOneConfig(config Config, client *client.Vproxy) ([]*Todo, error) {
	switch config.GetBase().Kind {
	case "TcpLb":
		fallthrough
	case "tcp-lb":
		fallthrough
	case "tl":
		return checkDeleteTcpLb(config, client)
	case "Socks5Server":
		fallthrough
	case "socks5-server":
		fallthrough
	case "socks5":
		return checkDeleteSocks5Server(config, client)
	case "DnsServer":
		fallthrough
	case "dns-server":
		fallthrough
	case "dns":
		return checkDeleteDnsServer(config, client)
	case "Upstream":
		fallthrough
	case "upstream":
		fallthrough
	case "ups":
		return checkDeleteUpstream(config, client)
	case "ServerGroup":
		fallthrough
	case "server-group":
		fallthrough
	case "sg":
		return checkDeleteServerGroup(config, client)
	case "SecurityGroup":
		fallthrough
	case "security-group":
		fallthrough
	case "secg":
		return checkDeleteSecurityGroup(config, client)
	case "CertKey":
		fallthrough
	case "cert-key":
		fallthrough
	case "ck":
		return checkDeleteCertKey(config, client)
	default:
		return nil, fmt.Errorf("BUG: unknown config %s", config)
	}
}

func checkDeleteTcpLb(config Config, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", config.GetBase().Kind, config.GetBase().Metadata.Name)
	name := config.GetBase().Metadata.Name
	_, err := GetTcpLb(name)
	if err != nil {
		if _, ok := err.(*tcp_lb.GetTCPLbNotFound); ok {
			// not found, so nothing to do
			return build404Op(ref)
		}
		return nil, err
	}
	return []*Todo{
		{
			client: client,
			op:     OpDelete,
			Ref:    ref,
			object: config,
			F:      deleteTcpLb,
		},
	}, nil
}

func checkDeleteSocks5Server(config Config, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", config.GetBase().Kind, config.GetBase().Metadata.Name)
	name := config.GetBase().Metadata.Name
	_, err := GetSocks5Server(name)
	if err != nil {
		if _, ok := err.(*socks5_server.GetSocks5ServerNotFound); ok {
			// not found, so nothing to do
			return build404Op(ref)
		}
		return nil, err
	}
	return []*Todo{
		{
			client: client,
			op:     OpDelete,
			Ref:    ref,
			object: config,
			F:      deleteSocks5Server,
		},
	}, nil
}

func checkDeleteDnsServer(config Config, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", config.GetBase().Kind, config.GetBase().Metadata.Name)
	name := config.GetBase().Metadata.Name
	_, err := GetDnsServer(name)
	if err != nil {
		if _, ok := err.(*dns_server.GetDNSServerNotFound); ok {
			// not found, so nothing to do
			return build404Op(ref)
		}
		return nil, err
	}
	return []*Todo{
		{
			client: client,
			op:     OpDelete,
			Ref:    ref,
			object: config,
			F:      deleteDnsServer,
		},
	}, nil
}

func checkDeleteUpstream(config Config, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", config.GetBase().Kind, config.GetBase().Metadata.Name)
	name := config.GetBase().Metadata.Name
	_, err := GetUpstream(name)
	if err != nil {
		if _, ok := err.(*upstream.GetUpstreamNotFound); ok {
			// not found, so nothing to do
			return build404Op(ref)
		}
		return nil, err
	}
	return []*Todo{
		{
			client: client,
			op:     OpDelete,
			Ref:    ref,
			object: config,
			F:      deleteUpstream,
		},
	}, nil
}

func checkDeleteServerGroup(config Config, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", config.GetBase().Kind, config.GetBase().Metadata.Name)
	name := config.GetBase().Metadata.Name
	_, err := GetServerGroup(name)
	if err != nil {
		if _, ok := err.(*server_group.GetServerGroupNotFound); ok {
			// not found, so nothing to do
			return build404Op(ref)
		}
		return nil, err
	}
	return []*Todo{
		{
			client: client,
			op:     OpDelete,
			Ref:    ref,
			object: config,
			F:      deleteServerGroup,
		},
	}, nil
}

func checkDeleteSecurityGroup(config Config, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", config.GetBase().Kind, config.GetBase().Metadata.Name)
	name := config.GetBase().Metadata.Name
	_, err := GetSecurityGroup(name)
	if err != nil {
		if _, ok := err.(*security_group.GetSecurityGroupNotFound); ok {
			// not found, so nothing to do
			return build404Op(ref)
		}
		return nil, err
	}
	return []*Todo{
		{
			client: client,
			op:     OpDelete,
			Ref:    ref,
			object: config,
			F:      deleteSecurityGroup,
		},
	}, nil
}

func checkDeleteCertKey(config Config, client *client.Vproxy) ([]*Todo, error) {
	ref := fmt.Sprintf("%s:%s", config.GetBase().Kind, config.GetBase().Metadata.Name)
	name := config.GetBase().Metadata.Name
	_, err := GetCertKey(name)
	if err != nil {
		if _, ok := err.(*cert_key.DescribeCertKeyNotFound); ok {
			// not found, so nothing to do
			return build404Op(ref)
		}
		return nil, err
	}
	return []*Todo{
		{
			client: client,
			op:     OpDelete,
			Ref:    ref,
			object: config,
			F:      deleteCertKey,
		},
	}, nil
}

func DeleteByConfig(configs []Config) ([]*Todo, error) {
	// validate
	for idx, c := range configs {
		err := c.Validate()
		if err != nil {
			return nil, fmt.Errorf("invalid document at index %d: %s", idx, err.Error())
		}
	}
	// sort
	tmp := ConfigSortingArrayForDeleting(configs)
	sort.Sort(tmp)
	configs = tmp
	// check
	cli := getClient()
	todoList := make([]*Todo, 0)
	for _, c := range configs {
		t, err := deleteOneConfig(c, cli)
		if err != nil {
			return nil, fmt.Errorf("cannot delete %s:%s: %s", c.GetBase().Kind, c.GetBase().Metadata.Name, err.Error())
		}
		todoList = append(todoList, t...)
	}
	return todoList, nil
}

func DeleteByFile(filename string) ([]*Todo, error) {
	configs, err := ParseFile(filename)
	if err != nil {
		return nil, err
	}
	return DeleteByConfig(configs)
}

func DeleteOne(kind string, name string) ([]*Todo, error) {
	configs := []Config{
		&Base{
			ApiVersion: CurrentVersion,
			Kind:       kind,
			Metadata: Metadata{
				Name: name,
			},
		},
	}
	return DeleteByConfig(configs)
}

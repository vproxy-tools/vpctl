package vproxy_config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"vproxy_client/cert_key"
	"vproxy_client/dns_server"
	"vproxy_client/security_group"
	"vproxy_client/security_group_rule"
	"vproxy_client/server"
	"vproxy_client/server_group"
	"vproxy_client/smart_group_delegate"
	"vproxy_client/socks5_server"
	"vproxy_client/tcp_lb"
	"vproxy_client/upstream"
	"vproxy_client_model"
)

func createTcpLb(todo *Todo) error {
	tl := todo.object.(*TcpLb)
	params := tcp_lb.NewAddTCPLbParams()
	params.SetBody(&vproxy_client_model.TCPLbCreate{
		Address:       &tl.Spec.Address,
		Backend:       &tl.Spec.Backend,
		ListOfCertKey: tl.Spec.ListOfCertKey,
		Name:          &tl.Metadata.Name,
		Protocol:      vproxy_client_model.Protocol(tl.Spec.Protocol),
		SecurityGroup: tl.Spec.SecurityGroup,
	})
	_, err := todo.client.TCPLb.AddTCPLb(params)
	return err
}

func updateTcpLb(todo *Todo) error {
	tl := todo.object.(*TcpLb)
	params := tcp_lb.NewUpdateTCPLbParams()
	params.SetBody(&vproxy_client_model.TCPLbUpdate{
		SecurityGroup: tl.Spec.SecurityGroup,
	})
	params.SetTl(tl.Metadata.Name)
	_, err := todo.client.TCPLb.UpdateTCPLb(params)
	return err
}

func deleteTcpLb(todo *Todo) error {
	tl := todo.object.(Config)
	params := tcp_lb.NewRemoveTCPLbParams()
	params.SetTl(tl.GetBase().Metadata.Name)
	_, err := todo.client.TCPLb.RemoveTCPLb(params)
	return err
}

func createSocks5Server(todo *Todo) error {
	socks5 := todo.object.(*Socks5Server)
	params := socks5_server.NewAddSocks5ServerParams()
	params.SetBody(&vproxy_client_model.Socks5ServerCreate{
		Address:         &socks5.Spec.Address,
		AllowNonBackend: socks5.Spec.AllowNonBackend,
		Backend:         &socks5.Spec.Backend,
		Name:            &socks5.Metadata.Name,
		SecurityGroup:   socks5.Spec.SecurityGroup,
	})
	_, err := todo.client.Socks5Server.AddSocks5Server(params)
	return err
}

func updateSocks5Server(todo *Todo) error {
	socks5 := todo.object.(*Socks5Server)
	params := socks5_server.NewUpdateSocks5ServerParams()
	params.SetBody(&vproxy_client_model.Socks5ServerUpdate{
		AllowNonBackend: socks5.Spec.AllowNonBackend,
		SecurityGroup:   socks5.Spec.SecurityGroup,
	})
	params.SetSocks5(socks5.Metadata.Name)
	_, err := todo.client.Socks5Server.UpdateSocks5Server(params)
	return err
}

func deleteSocks5Server(todo *Todo) error {
	socks5 := todo.object.(Config)
	params := socks5_server.NewRemoveSocks5ServerParams()
	params.SetSocks5(socks5.GetBase().Metadata.Name)
	_, err := todo.client.Socks5Server.RemoveSocks5Server(params)
	return err
}

func createDnsServer(todo *Todo) error {
	dns := todo.object.(*DnsServer)
	params := dns_server.NewAddDNSServerParams()
	params.SetBody(&vproxy_client_model.DNSServerCreate{
		Address:       &dns.Spec.Address,
		Name:          &dns.Metadata.Name,
		Rrsets:        &dns.Spec.RRSets,
		SecurityGroup: dns.Spec.SecurityGroup,
		TTL:           intptr2int64ptr(dns.Spec.TTL),
	})
	_, err := todo.client.DNSServer.AddDNSServer(params)
	return err
}

func updateDnsServer(todo *Todo) error {
	dns := todo.object.(*DnsServer)
	params := dns_server.NewUpdateDNSServerParams()
	params.SetBody(&vproxy_client_model.DNSServerUpdate{
		SecurityGroup: dns.Spec.SecurityGroup,
		TTL:           intptr2int64ptr(dns.Spec.TTL),
	})
	params.SetDNS(dns.Metadata.Name)
	_, err := todo.client.DNSServer.UpdateDNSServer(params)
	return err
}

func deleteDnsServer(todo *Todo) error {
	dns := todo.object.(Config)
	params := dns_server.NewRemoveDNSServerParams()
	params.SetDNS(dns.GetBase().Metadata.Name)
	_, err := todo.client.DNSServer.RemoveDNSServer(params)
	return err
}

func createUpstream(todo *Todo) error {
	ups := todo.object.(*Upstream)
	params := upstream.NewAddUpstreamParams()
	params.SetBody(&vproxy_client_model.UpstreamCreate{
		Name: &ups.Metadata.Name,
	})
	_, err := todo.client.Upstream.AddUpstream(params)
	return err
}

func updateUpstream(todo *Todo) error {
	_ = todo.object.(*Upstream)
	return fmt.Errorf("BUG: should not reach here: updateUpstream")
}

func deleteUpstream(todo *Todo) error {
	ups := todo.object.(Config)
	params := upstream.NewRemoveUpstreamParams()
	params.SetUps(ups.GetBase().Metadata.Name)
	_, err := todo.client.Upstream.RemoveUpstream(params)
	return err
}

func createServerGroupInUpstream(todo *Todo) error {
	ups := todo.parent.(*Upstream)
	sg := todo.object.(*ServerGroupInUpstream)
	params := server_group.NewAddServerGroupInUpstreamParams()
	params.SetBody(&vproxy_client_model.ServerGroupInUpstreamCreate{
		Annotations: sg.Annotations,
		Name:        &sg.Name,
		Weight:      intptr2int64ptr(sg.Weight),
	})
	params.SetUps(ups.Metadata.Name)
	_, err := todo.client.ServerGroup.AddServerGroupInUpstream(params)
	return err
}

func updateServerGroupInUpstream(todo *Todo) error {
	ups := todo.parent.(*Upstream)
	sg := todo.object.(*ServerGroupInUpstream)
	params := server_group.NewUpdateServerGroupInUpstreamParams()
	params.SetBody(&vproxy_client_model.ServerGroupInUpstreamUpdate{
		Annotations: sg.Annotations,
		Weight:      intptr2int64ptr(sg.Weight),
	})
	params.SetSg(sg.Name)
	params.SetUps(ups.Metadata.Name)
	_, err := todo.client.ServerGroup.UpdateServerGroupInUpstream(params)
	return err
}

func deleteServerGroupInUpstream(todo *Todo) error {
	ups := todo.parent.(*Upstream)
	sg := todo.object.(*ServerGroupInUpstream)
	params := server_group.NewRemoveServerGroupInUpstreamParams()
	params.SetSg(sg.Name)
	params.SetUps(ups.Metadata.Name)
	_, err := todo.client.ServerGroup.RemoveServerGroupInUpstream(params)
	return err
}

func createServerGroup(todo *Todo) error {
	sg := todo.object.(*ServerGroup)
	params := server_group.NewAddServerGroupParams()
	params.SetBody(&vproxy_client_model.ServerGroupCreate{
		Annotations: sg.Metadata.Annotations,
		Down:        int64ptr(sg.Spec.Down),
		Method:      vproxy_client_model.ServerGroupMethod(sg.Spec.Method),
		Name:        &sg.Metadata.Name,
		Period:      int64ptr(sg.Spec.Period),
		Protocol:    vproxy_client_model.CheckProtocol(sg.Spec.Protocol),
		Timeout:     int64ptr(sg.Spec.Timeout),
		Up:          int64ptr(sg.Spec.Up),
	})
	_, err := todo.client.ServerGroup.AddServerGroup(params)
	return err
}

func updateServerGroup(todo *Todo) error {
	sg := todo.object.(*ServerGroup)
	params := server_group.NewUpdateServerGroupParams()
	params.SetBody(&vproxy_client_model.ServerGroupUpdate{
		Annotations: sg.Metadata.Annotations,
		Down:        int64(sg.Spec.Down),
		Method:      vproxy_client_model.ServerGroupMethod(sg.Spec.Method),
		Period:      int64(sg.Spec.Period),
		Protocol:    vproxy_client_model.CheckProtocol(sg.Spec.Protocol),
		Timeout:     int64(sg.Spec.Timeout),
		Up:          int64(sg.Spec.Up),
	})
	params.SetSg(sg.Metadata.Name)
	_, err := todo.client.ServerGroup.UpdateServerGroup(params)
	return err
}

func deleteServerGroup(todo *Todo) error {
	sg := todo.object.(Config)
	params := server_group.NewRemoveServerGroupParams()
	params.SetSg(sg.GetBase().Metadata.Name)
	_, err := todo.client.ServerGroup.RemoveServerGroup(params)
	return err
}

func createStaticServerInServerGroup(todo *Todo) error {
	svr := todo.object.(*StaticServer)
	sg := todo.parent.(*ServerGroup)
	params := server.NewAddServerParams()
	params.SetBody(&vproxy_client_model.ServerCreate{
		Address: &svr.Address,
		Name:    &svr.Name,
		Weight:  intptr2int64ptr(svr.Weight),
	})
	params.SetSg(sg.Metadata.Name)
	_, err := todo.client.Server.AddServer(params)
	return err
}

func updateStaticServerInServerGroup(todo *Todo) error {
	svr := todo.object.(*StaticServer)
	sg := todo.parent.(*ServerGroup)
	params := server.NewUpdateServerParams()
	params.SetBody(&vproxy_client_model.ServerUpdate{
		Weight: intptr2int64ptr(svr.Weight),
	})
	params.SetSg(sg.Metadata.Name)
	params.SetSvr(svr.Name)
	_, err := todo.client.Server.UpdateServer(params)
	return err
}

func deleteStaticServerInServerGroup(todo *Todo) error {
	svr := todo.object.(*StaticServer)
	sg := todo.parent.(*ServerGroup)
	params := server.NewRemoveServerParams()
	params.SetSg(sg.Metadata.Name)
	params.SetSvr(svr.Name)
	_, err := todo.client.Server.RemoveServer(params)
	return err
}

func createSecurityGroup(todo *Todo) error {
	secg := todo.object.(*SecurityGroup)
	params := security_group.NewAddSecurityGroupParams()
	params.SetBody(&vproxy_client_model.SecurityGroupCreate{
		DefaultRule: vproxy_client_model.Rule(secg.Spec.DefaultRule),
		Name:        &secg.Metadata.Name,
	})
	_, err := todo.client.SecurityGroup.AddSecurityGroup(params)
	return err
}

func updateSecurityGroup(todo *Todo) error {
	secg := todo.object.(*SecurityGroup)
	params := security_group.NewUpdateSecurityGroupParams()
	params.SetBody(&vproxy_client_model.SecurityGroupUpdate{
		DefaultRule: vproxy_client_model.Rule(secg.Spec.DefaultRule),
	})
	params.SetSecg(secg.Metadata.Name)
	_, err := todo.client.SecurityGroup.UpdateSecurityGroup(params)
	return err
}

func deleteSecurityGroup(todo *Todo) error {
	secg := todo.object.(Config)
	params := security_group.NewRemoveSecurityGroupParams()
	params.SetSecg(secg.GetBase().Metadata.Name)
	_, err := todo.client.SecurityGroup.RemoveSecurityGroup(params)
	return err
}

func createSecurityGroupRuleInSecurityGroup(todo *Todo) error {
	secgr := todo.object.(*SecurityGroupRule)
	secg := todo.parent.(*SecurityGroup)
	params := security_group_rule.NewAddSecurityGroupRuleParams()
	params.SetBody(&vproxy_client_model.SecurityGroupRuleCreate{
		ClientNetwork: &secgr.ClientNetwork,
		Name:          &secgr.Name,
		Protocol:      vproxy_client_model.SecurityGroupProtocol(secgr.Protocol),
		Rule:          vproxy_client_model.Rule(secgr.Rule),
		ServerPortMax: int64ptr(secgr.ServerPortMax),
		ServerPortMin: int64ptr(secgr.ServerPortMin),
	})
	params.SetSecg(secg.Metadata.Name)
	_, err := todo.client.SecurityGroupRule.AddSecurityGroupRule(params)
	return err
}

func updateSecurityGroupRuleInSecurityGroup(todo *Todo) error {
	_ = todo.object.(*SecurityGroupRule)
	_ = todo.object.(*SecurityGroup)
	return fmt.Errorf("BUG: should not reach here: updateSecurityGroupRuleInSecurityGroup")
}

func deleteSecurityGroupRuleInSecurityGroup(todo *Todo) error {
	secgr := todo.object.(*SecurityGroupRule)
	secg := todo.parent.(*SecurityGroup)
	params := security_group_rule.NewRemoveSecurityGroupRuleParams()
	params.SetSecg(secg.Metadata.Name)
	params.SetSecgr(secgr.Name)
	_, err := todo.client.SecurityGroupRule.RemoveSecurityGroupRule(params)
	return err
}

func createCertKey(todo *Todo) error {
	ck := todo.object.(*CertKey)

	certNames := make([]string, 0)
	keyName := ""
	// write cert files
	for idx, content := range ck.Spec.Pem.Certs {
		filename := CertName(ck.Metadata.Name, idx)
		err := ioutil.WriteFile(filename, []byte(strings.TrimSpace(content)), 0644)
		if err != nil {
			return err
		}
		certNames = append(certNames, filename)
	}
	// write key file
	{
		filename := KeyName(ck.Metadata.Name)
		err := ioutil.WriteFile(filename, []byte(strings.TrimSpace(ck.Spec.Pem.Key)), 0644)
		if err != nil {
			return err
		}
		keyName = filename
	}

	params := cert_key.NewAddCertKeyParams()
	params.SetBody(&vproxy_client_model.CertKeyCreate{
		Certs: certNames,
		Key:   &keyName,
		Name:  &ck.Metadata.Name,
	})
	_, err := todo.client.CertKey.AddCertKey(params)
	return err
}

func updateCertKey(todo *Todo) error {
	_ = todo.object.(*CertKey)
	return fmt.Errorf("BUG: should not reach here: updateCertKey")
}

func deleteCertKey(todo *Todo) error {
	ck := todo.object.(Config)
	ckName := ck.GetBase().Metadata.Name
	workingDir := GetWorkingDir()
	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		fname := f.Name()
		if IsCertKeyName(fname, ckName) {
			fname = path.Join(workingDir, fname)
			err := os.Remove(fname)
			if err != nil {
				return err
			}
		}
	}

	params := cert_key.NewRemoveCertKeyParams()
	params.SetCk(ckName)
	_, err = todo.client.CertKey.RemoveCertKey(params)
	return err
}

func createSmartGroupDelegate(todo *Todo) error {
	sgd := todo.object.(*SmartGroupDelegate)
	params := smart_group_delegate.NewAddSmartGroupDelegateParams()
	params.SetBody(&vproxy_client_model.SmartGroupDelegateCreate{
		HandledGroup: &sgd.Spec.HandledGroup,
		Name:         &sgd.Metadata.Name,
		Service:      &sgd.Spec.Service,
		Zone:         &sgd.Spec.Zone,
	})
	_, err := todo.client.SmartGroupDelegate.AddSmartGroupDelegate(params)
	return err
}

func updateSmartGroupDelegate(todo *Todo) error {
	_ = todo.object.(*SmartGroupDelegate)
	return fmt.Errorf("BUG: should not reach here: updateSmartGroupDelegate")
}

func deleteSmartGroupDelegate(todo *Todo) error {
	sgd := todo.object.(Config)
	params := smart_group_delegate.NewRemoveSmartGroupDelegateParams()
	params.SetSgd(sgd.GetBase().Metadata.Name)
	_, err := todo.client.SmartGroupDelegate.RemoveSmartGroupDelegate(params)
	return err
}

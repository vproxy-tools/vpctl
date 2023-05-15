package vproxy_config

import (
	"fmt"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/cert_key"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/dns_server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/security_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/security_group_rule"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/server_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/socks5_server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/tcp_lb"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/upstream"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func createTcpLb(todo *Todo) error {
	tl := todo.object.(*TcpLb)
	params := tcp_lb.NewAddTCPLbParams()
	params.SetBody(func() *vproxy_client_model.TCPLbCreate { tmp := tl.toCreate(); return &tmp }())
	_, err := todo.client.TCPLb.AddTCPLb(params)
	return err
}

func updateTcpLb(todo *Todo) error {
	tl := todo.object.(*TcpLb)
	params := tcp_lb.NewUpdateTCPLbParams()
	params.SetBody(func() *vproxy_client_model.TCPLbUpdate { tmp := tl.toUpdate(); return &tmp }())
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
	params.SetBody(func() *vproxy_client_model.Socks5ServerCreate { tmp := socks5.toCreate(); return &tmp }())
	_, err := todo.client.Socks5Server.AddSocks5Server(params)
	return err
}

func updateSocks5Server(todo *Todo) error {
	socks5 := todo.object.(*Socks5Server)
	params := socks5_server.NewUpdateSocks5ServerParams()
	params.SetBody(func() *vproxy_client_model.Socks5ServerUpdate { tmp := socks5.toUpdate(); return &tmp }())
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
	params.SetBody(func() *vproxy_client_model.DNSServerCreate { tmp := dns.toCreate(); return &tmp }())
	_, err := todo.client.DNSServer.AddDNSServer(params)
	return err
}

func updateDnsServer(todo *Todo) error {
	dns := todo.object.(*DnsServer)
	params := dns_server.NewUpdateDNSServerParams()
	params.SetBody(func() *vproxy_client_model.DNSServerUpdate { tmp := dns.toUpdate(); return &tmp }())
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
	params.SetBody(func() *vproxy_client_model.UpstreamCreate { tmp := ups.toCreate(); return &tmp }())
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
	idx := todo.Index
	ups := todo.parent.(*Upstream)
	sg := todo.object.(*ServerGroupInUpstream)
	params := server_group.NewAddServerGroupInUpstreamParams()
	params.SetBody(func() *vproxy_client_model.ServerGroupInUpstreamCreate {
		tmp := sg.toCreate(ups.Metadata.Name, idx)
		return &tmp
	}())
	params.SetUps(ups.Metadata.Name)
	_, err := todo.client.ServerGroup.AddServerGroupInUpstream(params)
	return err
}

func updateServerGroupInUpstream(todo *Todo) error {
	idx := todo.Index
	ups := todo.parent.(*Upstream)
	sg := todo.object.(*ServerGroupInUpstream)
	params := server_group.NewUpdateServerGroupInUpstreamParams()
	params.SetBody(func() *vproxy_client_model.ServerGroupInUpstreamUpdate {
		tmp := sg.toUpdate(ups.Metadata.Name, idx)
		return &tmp
	}())
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
	params.SetBody(func() *vproxy_client_model.ServerGroupCreate { tmp := sg.toCreate(); return &tmp }())
	_, err := todo.client.ServerGroup.AddServerGroup(params)
	return err
}

func updateServerGroup(todo *Todo) error {
	sg := todo.object.(*ServerGroup)
	params := server_group.NewUpdateServerGroupParams()
	params.SetBody(func() *vproxy_client_model.ServerGroupUpdate { tmp := sg.toUpdate(); return &tmp }())
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
	idx := todo.Index
	svr := todo.object.(*StaticServer)
	sg := todo.parent.(*ServerGroup)
	params := server.NewAddServerParams()
	params.SetBody(func() *vproxy_client_model.ServerCreate { tmp := svr.toCreate(sg.Metadata.Name, idx); return &tmp }())
	params.SetSg(sg.Metadata.Name)
	_, err := todo.client.Server.AddServer(params)
	return err
}

func updateStaticServerInServerGroup(todo *Todo) error {
	idx := todo.Index
	svr := todo.object.(*StaticServer)
	sg := todo.parent.(*ServerGroup)
	params := server.NewUpdateServerParams()
	params.SetBody(func() *vproxy_client_model.ServerUpdate { tmp := svr.toUpdate(sg.Metadata.Name, idx); return &tmp }())
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
	params.SetBody(func() *vproxy_client_model.SecurityGroupCreate { tmp := secg.toCreate(); return &tmp }())
	_, err := todo.client.SecurityGroup.AddSecurityGroup(params)
	return err
}

func updateSecurityGroup(todo *Todo) error {
	secg := todo.object.(*SecurityGroup)
	params := security_group.NewUpdateSecurityGroupParams()
	params.SetBody(func() *vproxy_client_model.SecurityGroupUpdate { tmp := secg.toUpdate(); return &tmp }())
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
	idx := todo.Index
	secgr := todo.object.(*SecurityGroupRule)
	secg := todo.parent.(*SecurityGroup)
	params := security_group_rule.NewAddSecurityGroupRuleParams()
	params.SetBody(func() *vproxy_client_model.SecurityGroupRuleCreate {
		tmp := secgr.toCreate(secg.Metadata.Name, idx)
		return &tmp
	}())
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

	// write cert files
	for idx, content := range ck.Spec.Pem.Certs {
		filename := CertName(ck.Metadata.Name, idx)
		err := ioutil.WriteFile(filename, []byte(strings.TrimSpace(content)), 0644)
		if err != nil {
			return err
		}
	}
	// write key file
	{
		filename := KeyName(ck.Metadata.Name)
		err := ioutil.WriteFile(filename, []byte(strings.TrimSpace(ck.Spec.Pem.Key)), 0644)
		if err != nil {
			return err
		}
	}

	params := cert_key.NewAddCertKeyParams()
	params.SetBody(func() *vproxy_client_model.CertKeyCreate { tmp := ck.toCreate(); return &tmp }())
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

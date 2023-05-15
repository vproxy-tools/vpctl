package vproxy_config

import (
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/cert_key"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/dns_server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/security_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/security_group_rule"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/server_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/socks5_server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/tcp_lb"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/upstream"
)

func GetTcpLb(name string) (*TcpLb, error) {
	params := tcp_lb.NewGetTCPLbParams()
	params.SetTl(name)
	resp, err := getClient().TCPLb.GetTCPLb(params)
	if err != nil {
		return nil, err
	}
	tl := TcpLb{}
	tl.from(resp.GetPayload())
	return &tl, nil
}

func GetSocks5Server(name string) (*Socks5Server, error) {
	params := socks5_server.NewGetSocks5ServerParams()
	params.SetSocks5(name)
	resp, err := getClient().Socks5Server.GetSocks5Server(params)
	if err != nil {
		return nil, err
	}
	socks5 := Socks5Server{}
	socks5.from(resp.GetPayload())
	return &socks5, nil
}

func GetDnsServer(name string) (*DnsServer, error) {
	params := dns_server.NewGetDNSServerParams()
	params.SetDNS(name)
	resp, err := getClient().DNSServer.GetDNSServer(params)
	if err != nil {
		return nil, err
	}
	dns := DnsServer{}
	dns.from(resp.GetPayload())
	return &dns, nil
}

func GetUpstream(name string) (*Upstream, error) {
	params := upstream.NewGetUpstreamParams()
	params.SetUps(name)
	resp, err := getClient().Upstream.GetUpstream(params)
	if err != nil {
		return nil, err
	}
	params2 := server_group.NewListServerGroupInUpstreamParams()
	params2.SetUps(name)
	resp2, err := getClient().ServerGroup.ListServerGroupInUpstream(params2)
	if err != nil {
		return nil, err
	}
	ups := Upstream{}
	ups.from(resp.GetPayload(), resp2.GetPayload())
	return &ups, nil
}

func GetServerGroup(name string) (*ServerGroup, error) {
	params := server_group.NewGetServerGroupParams()
	params.SetSg(name)
	resp, err := getClient().ServerGroup.GetServerGroup(params)
	if err != nil {
		return nil, err
	}
	params2 := server.NewListServerParams()
	params2.SetSg(name)
	resp2, err := getClient().Server.ListServer(params2)
	if err != nil {
		return nil, err
	}
	sg := ServerGroup{}
	sg.from(resp.GetPayload(), resp2.GetPayload())
	return &sg, nil
}

func GetSecurityGroup(name string) (*SecurityGroup, error) {
	params := security_group.NewGetSecurityGroupParams()
	params.SetSecg(name)
	resp, err := getClient().SecurityGroup.GetSecurityGroup(params)
	if err != nil {
		return nil, err
	}
	params2 := security_group_rule.NewListSecurityGroupRuleParams()
	params2.SetSecg(name)
	resp2, err := getClient().SecurityGroupRule.ListSecurityGroupRule(params2)
	if err != nil {
		return nil, err
	}
	secg := SecurityGroup{}
	secg.from(resp.GetPayload(), resp2.GetPayload())
	return &secg, nil
}

func GetCertKey(name string) (*CertKey, error) {
	params := cert_key.NewDescribeCertKeyParams()
	params.SetCk(name)
	resp, err := getClient().CertKey.DescribeCertKey(params)
	if err != nil {
		return nil, err
	}
	ck := CertKey{}
	ck.from(resp.GetPayload())
	return &ck, nil
}

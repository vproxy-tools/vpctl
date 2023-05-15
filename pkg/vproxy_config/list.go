package vproxy_config

import (
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/cert_key"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/dns_server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/security_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/server_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/socks5_server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/tcp_lb"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/upstream"
)

func ListTcpLb() ([]*TcpLb, error) {
	params := tcp_lb.NewListTCPLbParams()
	resp, err := getClient().TCPLb.ListTCPLb(params)
	if err != nil {
		return nil, err
	}
	payloadList := resp.GetPayload()
	ret := make([]*TcpLb, len(payloadList))
	for idx, payload := range payloadList {
		tl := TcpLb{}
		tl.from(payload)
		ret[idx] = &tl
	}
	return ret, nil
}

func ListSocks5Server() ([]*Socks5Server, error) {
	params := socks5_server.NewListSocks5ServerParams()
	resp, err := getClient().Socks5Server.ListSocks5Server(params)
	if err != nil {
		return nil, err
	}
	payloadList := resp.GetPayload()
	ret := make([]*Socks5Server, len(payloadList))
	for idx, payload := range payloadList {
		socks5 := Socks5Server{}
		socks5.from(payload)
		ret[idx] = &socks5
	}
	return ret, nil
}

func ListDnsServer() ([]*DnsServer, error) {
	params := dns_server.NewListDNSServerParams()
	resp, err := getClient().DNSServer.ListDNSServer(params)
	if err != nil {
		return nil, err
	}
	payloadList := resp.GetPayload()
	ret := make([]*DnsServer, len(payloadList))
	for idx, payload := range payloadList {
		dns := DnsServer{}
		dns.from(payload)
		ret[idx] = &dns
	}
	return ret, nil
}

func ListUpstream() ([]*Upstream, error) {
	params := upstream.NewListUpstreamParams()
	resp, err := getClient().Upstream.ListUpstream(params)
	if err != nil {
		return nil, err
	}
	payloadList := resp.GetPayload()
	ret := make([]*Upstream, len(payloadList))
	for idx, payload := range payloadList {
		ups, err := GetUpstream(payload.Name)
		if err != nil {
			return nil, err
		}
		ret[idx] = ups
	}
	return ret, nil
}

func ListServerGroup() ([]*ServerGroup, error) {
	params := server_group.NewListServerGroupParams()
	resp, err := getClient().ServerGroup.ListServerGroup(params)
	if err != nil {
		return nil, err
	}
	payloadList := resp.GetPayload()
	ret := make([]*ServerGroup, len(payloadList))
	for idx, payload := range payloadList {
		sg, err := GetServerGroup(payload.Name)
		if err != nil {
			return nil, err
		}
		ret[idx] = sg
	}
	return ret, nil
}

func ListSecurityGroup() ([]*SecurityGroup, error) {
	params := security_group.NewListSecurityGroupParams()
	resp, err := getClient().SecurityGroup.ListSecurityGroup(params)
	if err != nil {
		return nil, err
	}
	payloadList := resp.GetPayload()
	ret := make([]*SecurityGroup, len(payloadList))
	for idx, payload := range payloadList {
		secg, err := GetSecurityGroup(payload.Name)
		if err != nil {
			return nil, err
		}
		ret[idx] = secg
	}
	return ret, nil
}

func ListCertKey() ([]*CertKey, error) {
	params := cert_key.NewListCertKeyParams()
	resp, err := getClient().CertKey.ListCertKey(params)
	if err != nil {
		return nil, err
	}
	payloadList := resp.GetPayload()
	ret := make([]*CertKey, len(payloadList))
	for idx, payload := range payloadList {
		ck, err := GetCertKey(payload.Name)
		if err != nil {
			return nil, err
		}
		ret[idx] = ck
	}
	return ret, nil
}

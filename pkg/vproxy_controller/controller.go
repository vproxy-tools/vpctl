package vproxy_controller

import (
	"crypto/x509"
)

type WatchConfig struct {
	UseHttps bool
	HostPort string
	Token    string
	CaRoots  *x509.CertPool
	Labels   map[string]string

	lastVer string
}

type meta struct {
	ns   string
	name string
	ver  int
}

func Launch(conf WatchConfig) {
	pool := NewPool()
	pool.RegisterEventHandler(NewVpctl(conf.Labels))
	go watchTcpLb(&conf, pool)
	go watchSocks5Server(&conf, pool)
	go watchDnsServer(&conf, pool)
	go watchUpstream(&conf, pool)
	go watchServerGroup(&conf, pool)
	go watchSecurityGroup(&conf, pool)
	go watchEndpoints(&conf, pool)
	go watchSecrets(&conf, pool)
}

func watchTcpLb(conf *WatchConfig, pool Pool) {
	watch("TcpLb", conf, "/apis/vproxy.cc/v1alpha1/tcp-lb",
		DefineTcpLbFunc(pool), DeleteTcpLbFunc(pool), ClearTcpLbFunc(pool))
}

func watchSocks5Server(conf *WatchConfig, pool Pool) {
	watch("Socks5Server", conf, "/apis/vproxy.cc/v1alpha1/socks5-server",
		DefineSocks5ServerFunc(pool), DeleteSocks5ServerFunc(pool), ClearSocks5ServerFunc(pool))
}

func watchDnsServer(conf *WatchConfig, pool Pool) {
	watch("DnsServer", conf, "/apis/vproxy.cc/v1alpha1/dns-server",
		DefineDnsServerFunc(pool), DeleteDnsServerFunc(pool), ClearDnsServerFunc(pool))
}

func watchUpstream(conf *WatchConfig, pool Pool) {
	watch("Upstream", conf, "/apis/vproxy.cc/v1alpha1/upstream",
		DefineUpstreamFunc(pool), DeleteUpstreamFunc(pool), ClearUpstreamFunc(pool))
}

func watchServerGroup(conf *WatchConfig, pool Pool) {
	watch("ServerGroup", conf, "/apis/vproxy.cc/v1alpha1/server-group",
		DefineServerGroupFunc(pool), DeleteServerGroupFunc(pool), ClearServerGroupFunc(pool))
}

func watchSecurityGroup(conf *WatchConfig, pool Pool) {
	watch("ServerGroup", conf, "/apis/vproxy.cc/v1alpha1/security-group",
		DefineSecurityGroupFunc(pool), DeleteSecurityGroupFunc(pool), ClearSecurityGroupFunc(pool))
}

func watchEndpoints(conf *WatchConfig, pool Pool) {
	watch("Endpoints", conf, "/api/v1/endpoints",
		DefineEndpointsFunc(pool), DeleteEndpointsFunc(pool), ClearEndpointsFunc(pool))
}

func watchSecrets(conf *WatchConfig, pool Pool) {
	watch("Secret", conf, "/api/v1/secrets",
		DefineSecretFunc(pool), DeleteSecretFunc(pool), ClearSecretFunc(pool))
}

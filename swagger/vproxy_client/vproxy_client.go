// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/cert_key"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/dns_server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/event_loop"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/event_loop_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/security_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/security_group_rule"
	serverops "github.com/vproxy-tools/vpctl/swagger/vproxy_client/server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/server_group"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/socks5_server"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/tcp_lb"
	"github.com/vproxy-tools/vpctl/swagger/vproxy_client/upstream"
)

// Default vproxy HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "127.0.0.1:18776"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1/module"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new vproxy HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Vproxy {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new vproxy HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Vproxy {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new vproxy client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Vproxy {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Vproxy)
	cli.Transport = transport
	cli.CertKey = cert_key.New(transport, formats)
	cli.DNSServer = dns_server.New(transport, formats)
	cli.EventLoop = event_loop.New(transport, formats)
	cli.EventLoopGroup = event_loop_group.New(transport, formats)
	cli.SecurityGroup = security_group.New(transport, formats)
	cli.SecurityGroupRule = security_group_rule.New(transport, formats)
	cli.Server = serverops.New(transport, formats)
	cli.ServerGroup = server_group.New(transport, formats)
	cli.Socks5Server = socks5_server.New(transport, formats)
	cli.TCPLb = tcp_lb.New(transport, formats)
	cli.Upstream = upstream.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Vproxy is a client for vproxy
type Vproxy struct {
	CertKey cert_key.ClientService

	DNSServer dns_server.ClientService

	EventLoop event_loop.ClientService

	EventLoopGroup event_loop_group.ClientService

	SecurityGroup security_group.ClientService

	SecurityGroupRule security_group_rule.ClientService

	Server serverops.ClientService

	ServerGroup server_group.ClientService

	Socks5Server socks5_server.ClientService

	TCPLb tcp_lb.ClientService

	Upstream upstream.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Vproxy) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.CertKey.SetTransport(transport)
	c.DNSServer.SetTransport(transport)
	c.EventLoop.SetTransport(transport)
	c.EventLoopGroup.SetTransport(transport)
	c.SecurityGroup.SetTransport(transport)
	c.SecurityGroupRule.SetTransport(transport)
	c.Server.SetTransport(transport)
	c.ServerGroup.SetTransport(transport)
	c.Socks5Server.SetTransport(transport)
	c.TCPLb.SetTransport(transport)
	c.Upstream.SetTransport(transport)
}
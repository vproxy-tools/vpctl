package vproxy_controller

import (
	"../vproxy_config"
)

type ChunkMetadata struct {
	ResourceVersion string `json:"resourceVersion" yaml:"resourceVersion"`
	Namespace       string `json:"namespace" yaml:"namespace"`
	Name            string `json:"name" yaml:"name"`
}

type ChunkObject struct {
	Kind     string        `json:"kind" yaml:"kind"`
	Metadata ChunkMetadata `json:"metadata" yaml:"metadata"`
}

type Chunk struct {
	Type   string      `json:"type" yaml:"type"`
	Object ChunkObject `json:"object" yaml:"object"`
}

type Selector map[string]string

type Meta struct {
	Pending bool // the resource watcher is re-syncing
	Version int  // the version of latest related event
}

type TcpLb struct {
	vproxy_config.Base `yaml:",inline"`
	Selector           Selector                `json:"selector" yaml:"selector"`
	Spec               vproxy_config.TcpLbSpec `json:"spec" yaml:"spec"`
	M                  Meta                    `json:"-"`
}

type Socks5Server struct {
	vproxy_config.Base `yaml:",inline"`
	Selector           Selector                       `json:"selector" yaml:"selector"`
	Spec               vproxy_config.Socks5ServerSpec `json:"spec" yaml:"spec"`
	M                  Meta                           `json:"-"`
}

type DnsServer struct {
	vproxy_config.Base `yaml:",inline"`
	Selector           Selector                    `json:"selector" yaml:"selector"`
	Spec               vproxy_config.DnsServerSpec `json:"spec" yaml:"spec"`
	M                  Meta                        `json:"-"`
}

type Upstream struct {
	vproxy_config.Upstream `yaml:",inline"`
	M                      Meta `json:"-"`
}

type SecurityGroup struct {
	vproxy_config.SecurityGroup `yaml:",inline"`
	M                           Meta `json:"-"`
}

type ServiceRef struct {
	Name   string `json:"name" yaml:"name"`
	Weight *int   `json:"weight" yaml:"weight"`
}

type ServerInServerGroup struct {
	Services []*ServiceRef `json:"services" yaml:"services"`
}

type ServerGroupSpec struct {
	vproxy_config.ServerGroupSelfSpec `yaml:",inline"`
	Servers                           ServerInServerGroup `json:"servers" yaml:"servers"`
}

type ServerGroup struct {
	vproxy_config.Base `yaml:",inline"`
	Spec               ServerGroupSpec `json:"spec" yaml:"spec"`
	M                  Meta            `json:"-"`
}

type EndpointsAddress struct {
	Hostname string `json:"hostname" yaml:"hostname"`
	Ip       string `json:"ip" yaml:"ip"`
}

type EndpointsPort struct {
	Port int `json:"port" yaml:"port"`
}

type EndpointsSubset struct {
	Addresses []EndpointsAddress `json:"addresses" yaml:"addresses"`
	Ports     []EndpointsPort    `json:"ports" yaml:"ports"`
}

type Endpoints struct {
	vproxy_config.Base `yaml:",inline"`
	Subsets            []EndpointsSubset `json:"subsets" yaml:"subsets"`
	M                  Meta              `json:"-"`
}

type SecretData struct {
	Crt string `json:"tls.crt" yaml:"tls.crt"`
	Key string `json:"tls.key" yaml:"tls.key"`
}

type Secret struct {
	vproxy_config.Base `yaml:",inline"`
	Data               SecretData `json:"data" yaml:"data"`
	Type               string     `json:"type" yaml:"type"`
	M                  Meta       `json:"-"`
}

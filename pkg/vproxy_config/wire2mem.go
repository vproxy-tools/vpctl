package vproxy_config

// transform from http get request to entity

import (
	"strings"
	model "vproxy_client_model"
)

func (o *TcpLb) from(x *model.TCPLb) {
	o.ApiVersion = CurrentVersion
	o.Kind = "TcpLb"
	o.Metadata.Name = x.Name
	o.Spec.Address = x.Address
	o.Spec.Backend = x.Backend
	o.Spec.Protocol = string(x.Protocol)
	o.Spec.ListOfCertKey = x.ListOfCertKey
	o.Spec.SecurityGroup = x.SecurityGroup
}

func (o *Socks5Server) from(x *model.Socks5Server) {
	o.ApiVersion = CurrentVersion
	o.Kind = "Socks5Server"
	o.Metadata.Name = x.Name
	o.Spec.Address = x.Address
	o.Spec.Backend = x.Backend
	o.Spec.SecurityGroup = x.SecurityGroup
	o.Spec.AllowNonBackend = &x.AllowNonBackend
}

func (o *DnsServer) from(x *model.DNSServer) {
	o.ApiVersion = CurrentVersion
	o.Kind = "DnsServer"
	o.Metadata.Name = x.Name
	o.Spec.Address = x.Address
	o.Spec.RRSets = x.Rrsets
	o.Spec.TTL = int642intptr(x.TTL)
	o.Spec.SecurityGroup = x.SecurityGroup
}

func (o *Upstream) from(x *model.Upstream, y []*model.ServerGroupInUpstream) {
	o.ApiVersion = CurrentVersion
	o.Kind = "Upstream"
	o.Metadata.Name = x.Name
	o.Spec.ServerGroups = make([]ServerGroupInUpstream, len(y))
	for idx, z := range y {
		a := ServerGroupInUpstream{
			Name:        z.Name,
			Weight:      int642intptr(z.Weight),
			Annotations: z.Annotations,
		}
		o.Spec.ServerGroups[idx] = a
	}
}

func (o *ServerGroup) from(x *model.ServerGroup, y []*model.Server) {
	o.ApiVersion = CurrentVersion
	o.Kind = "ServerGroup"
	o.Metadata.Name = x.Name
	o.Metadata.Annotations = x.Annotations
	o.Spec.Timeout = int(x.Timeout)
	o.Spec.Period = int(x.Period)
	o.Spec.Up = int(x.Up)
	o.Spec.Down = int(x.Down)
	o.Spec.Protocol = string(x.Protocol)
	o.Spec.Method = string(x.Method)
	o.Spec.Servers.Static = make([]StaticServer, len(y))
	for idx, z := range y {
		a := StaticServer{
			Name:    z.Name,
			Address: z.Address,
			Weight:  int642intptr(z.Weight),
		}
		o.Spec.Servers.Static[idx] = a
	}
	o.Status.Servers = make([]ServerStatus, len(y))
	for idx, z := range y {
		a := ServerStatus{
			Name:       z.Name,
			Address:    z.Address,
			Weight:     int(z.Weight),
			CurrentIp:  z.CurrentIP,
			Status:     z.Status,
			Cost:       int(z.Cost),
			DownReason: z.DownReason,
		}
		o.Status.Servers[idx] = a
	}
}

func (o *SecurityGroup) from(x *model.SecurityGroup, y []*model.SecurityGroupRule) {
	o.ApiVersion = CurrentVersion
	o.Kind = "SecurityGroup"
	o.Metadata.Name = x.Name
	o.Spec.DefaultRule = string(x.DefaultRule)
	o.Spec.Rules = make([]SecurityGroupRule, len(y))
	for idx, z := range y {
		a := SecurityGroupRule{
			Name:          z.Name,
			ClientNetwork: z.ClientNetwork,
			Protocol:      string(z.Protocol),
			ServerPortMin: int(z.ServerPortMin),
			ServerPortMax: int(z.ServerPortMax),
			Rule:          string(z.Rule),
		}
		o.Spec.Rules[idx] = a
	}
}

func (o *CertKey) from(x *model.CertKeyDetail) {
	o.ApiVersion = CurrentVersion
	o.Kind = "CertKey"
	o.Metadata.Name = x.Name
	o.Spec.Pem.Certs = make([]string, len(x.Certs))
	for idx, z := range x.CertPemList {
		o.Spec.Pem.Certs[idx] = strings.TrimSpace(z)
	}
	o.Spec.Pem.Key = "(hidden)"
	o.Status.CertFiles = make([]string, len(x.Certs))
	for idx, z := range x.Certs {
		o.Status.CertFiles[idx] = z
	}
	o.Status.KeyFile = x.Key
	o.Status.KeySHA1 = x.KeySHA1
}

func (o *HealthCheckEvent) from(x *HealthCheckEventOnWire) {
	sg := ServerGroup{}
	sg.from(&x.ServerGroup, []*model.Server{&x.Server})

	o.ApiVersion = CurrentVersion
	o.Kind = "HealthCheck"
	o.Metadata.Name = "event"
	o.Server = sg.Status.Servers[0]
	o.ServerGroup = ServerGroupInHC{
		Base: sg.Base,
		Spec: sg.Spec.ServerGroupSelfSpec,
	}
}

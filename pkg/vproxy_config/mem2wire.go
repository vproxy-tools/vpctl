package vproxy_config

// transform from entity to http request

import (
	model "github.com/vproxy-tools/vpctl/swagger/vproxy_client_model"
)

func (o *TcpLb) toCreate() model.TCPLbCreate {
	return model.TCPLbCreate{
		Address:       &o.Spec.Address,
		Backend:       &o.Spec.Backend,
		ListOfCertKey: o.Spec.ListOfCertKey,
		Name:          &o.Metadata.Name,
		Protocol:      model.Protocol(o.Spec.Protocol),
		SecurityGroup: o.Spec.SecurityGroup,
	}
}

func (o *TcpLb) toUpdate() model.TCPLbUpdate {
	return model.TCPLbUpdate{
		ListOfCertKey: o.Spec.ListOfCertKey,
		SecurityGroup: o.Spec.SecurityGroup,
	}
}

func (o *Socks5Server) toCreate() model.Socks5ServerCreate {
	return model.Socks5ServerCreate{
		Address:         &o.Spec.Address,
		AllowNonBackend: o.Spec.AllowNonBackend,
		Backend:         &o.Spec.Backend,
		Name:            &o.Metadata.Name,
		SecurityGroup:   o.Spec.SecurityGroup,
	}
}

func (o *Socks5Server) toUpdate() model.Socks5ServerUpdate {
	return model.Socks5ServerUpdate{
		AllowNonBackend: o.Spec.AllowNonBackend,
		SecurityGroup:   o.Spec.SecurityGroup,
	}
}

func (o *DnsServer) toCreate() model.DNSServerCreate {
	return model.DNSServerCreate{
		Address:       &o.Spec.Address,
		Name:          &o.Metadata.Name,
		Rrsets:        &o.Spec.RRSets,
		SecurityGroup: o.Spec.SecurityGroup,
		TTL:           intptr2int64ptr(o.Spec.TTL),
	}
}

func (o *DnsServer) toUpdate() model.DNSServerUpdate {
	return model.DNSServerUpdate{
		SecurityGroup: o.Spec.SecurityGroup,
		TTL:           intptr2int64ptr(o.Spec.TTL),
	}
}

func (o *StaticServer) toCreate(parent string, idx int) model.ServerCreate {
	return model.ServerCreate{
		Address: &o.Address,
		Name:    &o.Name,
		Weight:  intptr2int64ptr(o.Weight),
	}
}

func (o *StaticServer) toUpdate(parent string, idx int) model.ServerUpdate {
	return model.ServerUpdate{
		Weight: intptr2int64ptr(o.Weight),
	}
}

func int64ptr(i int) *int64 {
	x := int64(i)
	return &x
}

func intptr2int64ptr(i *int) *int64 {
	if i == nil {
		return nil
	}
	return int64ptr(*i)
}

func int642intptr(i int64) *int {
	x := int(i)
	return &x
}

func (o *ServerGroup) toCreate() model.ServerGroupCreate {
	return model.ServerGroupCreate{
		Annotations: o.Metadata.Annotations,
		Down:        int64ptr(o.Spec.Down),
		Method: func() *model.ServerGroupMethod {
			if o.Spec.Method == "" {
				return nil
			}
			tmp := model.ServerGroupMethod(o.Spec.Method)
			return &tmp
		}(),
		Name:   &o.Metadata.Name,
		Period: int64ptr(o.Spec.Period),
		Protocol: func() *model.CheckProtocol {
			if o.Spec.Protocol == "" {
				return nil
			}
			tmp := model.CheckProtocol(o.Spec.Protocol)
			return &tmp
		}(),
		Timeout: int64ptr(o.Spec.Timeout),
		Up:      int64ptr(o.Spec.Up),
	}
}

func (o *ServerGroup) toUpdate() model.ServerGroupUpdate {
	return model.ServerGroupUpdate{
		Annotations: o.Metadata.Annotations,
		Down:        int64(o.Spec.Down),
		Method: func() *model.ServerGroupMethod {
			if o.Spec.Method == "" {
				return nil
			}
			tmp := model.ServerGroupMethod(o.Spec.Method)
			return &tmp
		}(),
		Period: int64(o.Spec.Period),
		Protocol: func() *model.CheckProtocol {
			if o.Spec.Protocol == "" {
				return nil
			}
			tmp := model.CheckProtocol(o.Spec.Protocol)
			return &tmp
		}(),
		Timeout: int64(o.Spec.Timeout),
		Up:      int64(o.Spec.Up),
	}
}

func (o *ServerGroupInUpstream) toCreate(parent string, idx int) model.ServerGroupInUpstreamCreate {
	return model.ServerGroupInUpstreamCreate{
		Annotations: o.Annotations,
		Name:        &o.Name,
		Weight:      intptr2int64ptr(o.Weight),
	}
}

func (o *ServerGroupInUpstream) toUpdate(parent string, idx int) model.ServerGroupInUpstreamUpdate {
	return model.ServerGroupInUpstreamUpdate{
		Annotations: o.Annotations,
		Weight:      intptr2int64ptr(o.Weight),
	}
}

func (o *Upstream) toCreate() model.UpstreamCreate {
	return model.UpstreamCreate{
		Name: &o.Metadata.Name,
	}
}

func (o *SecurityGroupRule) toCreate(parent string, idx int) model.SecurityGroupRuleCreate {
	return model.SecurityGroupRuleCreate{
		ClientNetwork: &o.ClientNetwork,
		Name:          &o.Name,
		Protocol: func() *model.SecurityGroupProtocol {
			if o.Protocol == "" {
				return nil
			}
			tmp := model.SecurityGroupProtocol(o.Protocol)
			return &tmp
		}(),
		Rule: func() *model.Rule {
			if o.Rule == "" {
				return nil
			}
			tmp := model.Rule(o.Rule)
			return &tmp
		}(),
		ServerPortMax: int64ptr(o.ServerPortMax),
		ServerPortMin: int64ptr(o.ServerPortMin),
	}
}

func (o *SecurityGroup) toCreate() model.SecurityGroupCreate {
	return model.SecurityGroupCreate{
		DefaultRule: func() *model.Rule {
			if o.Spec.DefaultRule == "" {
				return nil
			}
			tmp := model.Rule(o.Spec.DefaultRule)
			return &tmp
		}(),
		Name: &o.Metadata.Name,
	}
}

func (o *SecurityGroup) toUpdate() model.SecurityGroupUpdate {
	return model.SecurityGroupUpdate{
		DefaultRule: model.Rule(o.Spec.DefaultRule),
	}
}

func (o *CertKey) toCreate() model.CertKeyCreate {
	certFiles := make([]string, len(o.Spec.Pem.Certs))
	for idx := range o.Spec.Pem.Certs {
		certFiles[idx] = CertName(o.Metadata.Name, idx)
	}
	keyFile := KeyName(o.Metadata.Name)
	return model.CertKeyCreate{
		Certs: certFiles,
		Key:   &keyFile,
		Name:  &o.Metadata.Name,
	}
}

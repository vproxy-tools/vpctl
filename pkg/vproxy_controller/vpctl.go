package vproxy_controller

import (
	"../vproxy_config"
	"encoding/base64"
	"strconv"
)

func NewVpctl(labels map[string]string) *Vpctl {
	return &Vpctl{
		Labels: labels,
	}
}

type Vpctl struct {
	Labels map[string]string
}

func (a *Vpctl) Name() string {
	return "vpctl"
}

func (a *Vpctl) Handle(pool *NamespacedResourcePool) (err error) {
	Debug("vpctl.handle(%v)", pool)

	configs := a.generate(pool)

	Debug("generated configs: %v", configs)

	deleteTodoList := make([]*vproxy_config.Todo, 0)

	var todo []*vproxy_config.Todo

	todo, err = GcTcpLb(configs)
	if err != nil {
		return
	}
	deleteTodoList = append(deleteTodoList, todo...)

	todo, err = GcSocks5Server(configs)
	if err != nil {
		return
	}
	deleteTodoList = append(deleteTodoList, todo...)

	todo, err = GcDnsServer(configs)
	if err != nil {
		return
	}
	deleteTodoList = append(deleteTodoList, todo...)

	todo, err = GcUpstream(configs)
	if err != nil {
		return
	}
	deleteTodoList = append(deleteTodoList, todo...)

	todo, err = GcServerGroup(configs)
	if err != nil {
		return
	}
	deleteTodoList = append(deleteTodoList, todo...)

	todo, err = GcSecurityGroup(configs)
	if err != nil {
		return
	}
	deleteTodoList = append(deleteTodoList, todo...)

	todo, err = GcCertKey(configs)
	if err != nil {
		return
	}
	deleteTodoList = append(deleteTodoList, todo...)

	applyTodoList, err := vproxy_config.ApplyByConfig(configs)
	if err != nil {
		return
	}

	allTodoList := merge(deleteTodoList, applyTodoList)

	hasChanges := false
	for _, task := range allTodoList {
		if task.Op != vproxy_config.OpText && task.Op != vproxy_config.OpNone && task.Op != vproxy_config.Op404 {
			hasChanges = true
			break
		}
	}
	if !hasChanges {
		return
	}
	for _, task := range allTodoList {
		if task.Op == vproxy_config.OpText || task.Op == vproxy_config.OpNone {
			continue
		}
		if err = task.F(task); err != nil {
			if task.Op == vproxy_config.OpCreate {
				Log("creating %v failed: %v", task.Ref, err.Error())
			} else if task.Op == vproxy_config.OpDelete {
				Log("deleting %v failed: %v", task.Ref, err.Error())
			} else {
				Log("updating %v failed: %v", task.Ref, err.Error())
			}
			return
		} else {
			if task.Op == vproxy_config.OpCreate {
				Log(task.Ref + " created")
			} else if task.Op == vproxy_config.OpDelete {
				Log(task.Ref + " deleted")
			} else {
				Log(task.Ref + " configured")
			}
		}
	}
	return nil
}

func (v *Vpctl) filterLabels(selector Selector) bool {
	for key, sVal := range selector {
		if lVal, exists := v.Labels[key]; !exists || lVal != sVal {
			return false
		}
	}
	return true
}

func findEpByName(eps []*Endpoints, n string) *Endpoints {
	for _, ep := range eps {
		if ep.Metadata.Name == n {
			return ep
		}
	}
	return nil
}

func makeStaticServers(ns string, s *ServerGroup, eps []*Endpoints) []vproxy_config.StaticServer {
	addrMap := map[string][]string{}
	weightMap := map[string]int{}
	for _, ref := range s.Spec.Servers.Services {
		if ref.Weight == nil {
			ref.Weight = int2ptr(10)
		}
		if *ref.Weight == 0 {
			continue // ignore the subset with weight 0
		}
		name := ref.Name
		ep := findEpByName(eps, name)
		if ep == nil {
			continue // ignore the subset if ep not found
		}
		l4addrLs := make([]string, 0)
		for _, subset := range ep.Subsets {
			for _, port := range subset.Ports {
				for _, addr := range subset.Addresses {
					if addr.Ip != "" {
						l4addrLs = append(l4addrLs, addr.Ip+":"+strconv.Itoa(port.Port))
					} else if addr.Hostname != "" {
						l4addrLs = append(l4addrLs, addr.Hostname+":"+strconv.Itoa(port.Port))
					} else {
						Log("The resource (Endpoints %s/%s) has an endpoint without ip nor hostname, ignore", ns, ep.Metadata.Name)
					}
				}
			}
		}
		if len(l4addrLs) == 0 {
			continue // ignore the subset if ep is empty
		}
		addrMap[name] = l4addrLs
		weightMap[name] = *ref.Weight
	}
	if len(addrMap) == 0 {
		return make([]vproxy_config.StaticServer, 0)
	}
	if len(addrMap) > 1 {
		for name := range addrMap {
			otherCount := 0
			for name2, addrs2 := range addrMap {
				if name == name2 {
					continue
				}
				otherCount += len(addrs2)
			}
			weightMap[name] = weightMap[name] * otherCount
		}
	}
	ngcd := -1
	for _, weight := range weightMap {
		if ngcd == -1 {
			ngcd = weight
		} else {
			ngcd = gcd(ngcd, weight)
		}
	}
	for name, weight := range weightMap {
		weightMap[name] = weight / ngcd
	}
	ret := make([]vproxy_config.StaticServer, 0)
	for name, l4addrs := range addrMap {
		for idx, l4addr := range l4addrs {
			svr := vproxy_config.StaticServer{
				Name:    makeName(ns, s.Metadata.Name) + "_" + strconv.Itoa(idx) + "_" + name,
				Address: l4addr,
				Weight:  int2ptr(weightMap[name]),
			}
			ret = append(ret, svr)
		}
	}
	return ret
}

func makeName(ns string, name string) string {
	if name == "" { // "" means name not provided, return default value
		return ""
	}
	return ns + "_" + name
}

func makeBase(ns string, kind string, name string) vproxy_config.Base {
	return vproxy_config.Base{
		ApiVersion: vproxy_config.CurrentVersion,
		Kind:       kind,
		Metadata: vproxy_config.Metadata{
			Name: makeName(ns, name),
		},
	}
}

func makeBaseWithMetadata(ns string, kind string, meta vproxy_config.Metadata) vproxy_config.Base {
	return vproxy_config.Base{
		ApiVersion: vproxy_config.CurrentVersion,
		Kind:       kind,
		Metadata: vproxy_config.Metadata{
			Name:        makeName(ns, meta.Name),
			Annotations: meta.Annotations,
		},
	}
}

func makeBaseWithBase(ns string, base vproxy_config.Base) vproxy_config.Base {
	return makeBaseWithMetadata(ns, base.Kind, base.Metadata)
}

func (v *Vpctl) generate(pool *NamespacedResourcePool) []vproxy_config.Config {
	ret := make([]vproxy_config.Config, 0)
	for ns, p := range pool.pools {
		upsInUse := map[string]bool{}
		secgInUse := map[string]bool{}
		ckInUse := map[string]bool{}
		sgInUse := map[string]bool{}

		foundUps := map[string]bool{}
		foundSecg := map[string]bool{}
		foundSg := map[string]bool{}

	loopTl:
		for _, t := range p.tl {
			if !v.filterLabels(t.Selector) {
				continue
			}
			if len(t.Spec.ListOfCertKey) != 0 {
				// we cannot make an empty ck that does not exist, so check whether the ck is provided
				for _, x := range t.Spec.ListOfCertKey {
					found := false
					for _, y := range p.ck {
						if y.Metadata.Name == x {
							found = true
							break
						}
					}
					if !found {
						Log("cert-key %v is not found in namespace %v, so tcp-lb %v cannot be created", x, ns, t.Metadata.Name)
						continue loopTl
					}
				}
			}

			listOfCertKey := make([]string, len(t.Spec.ListOfCertKey))
			for i, s := range t.Spec.ListOfCertKey {
				listOfCertKey[i] = makeName(ns, s)
			}
			vt := vproxy_config.TcpLb{
				Base: makeBaseWithBase(ns, t.Base),
				Spec: vproxy_config.TcpLbSpec{
					Address:       t.Spec.Address,
					Backend:       makeName(ns, t.Spec.Backend),
					Protocol:      t.Spec.Protocol,
					ListOfCertKey: listOfCertKey,
					SecurityGroup: makeName(ns, t.Spec.SecurityGroup),
				},
			}
			upsInUse[t.Spec.Backend] = true
			if t.Spec.SecurityGroup != "" {
				secgInUse[t.Spec.SecurityGroup] = true
			}
			if len(t.Spec.ListOfCertKey) != 0 {
				for _, x := range t.Spec.ListOfCertKey {
					ckInUse[x] = true
				}
			}
			ret = append(ret, &vt)
		}
		for _, s := range p.socks5 {
			if !v.filterLabels(s.Selector) {
				continue
			}
			vs := vproxy_config.Socks5Server{
				Base: makeBaseWithBase(ns, s.Base),
				Spec: vproxy_config.Socks5ServerSpec{
					Address:         s.Spec.Address,
					Backend:         makeName(ns, s.Spec.Backend),
					SecurityGroup:   makeName(ns, s.Spec.SecurityGroup),
					AllowNonBackend: nil,
				},
			}
			upsInUse[s.Spec.Backend] = true
			if s.Spec.SecurityGroup != "" {
				secgInUse[s.Spec.SecurityGroup] = true
			}
			ret = append(ret, &vs)
		}
		for _, d := range p.dns {
			if !v.filterLabels(d.Selector) {
				continue
			}
			vd := vproxy_config.DnsServer{
				Base: makeBaseWithBase(ns, d.Base),
				Spec: vproxy_config.DnsServerSpec{
					Address:       d.Spec.Address,
					RRSets:        makeName(ns, d.Spec.RRSets),
					TTL:           d.Spec.TTL,
					SecurityGroup: makeName(ns, d.Spec.SecurityGroup),
				},
			}
			upsInUse[d.Spec.RRSets] = true
			if d.Spec.SecurityGroup != "" {
				secgInUse[d.Spec.SecurityGroup] = true
			}
			ret = append(ret, &vd)
		}
		for _, u := range p.ups {
			if !upsInUse[u.Metadata.Name] {
				continue
			}
			foundUps[u.Metadata.Name] = true
			serverGroups := make([]vproxy_config.ServerGroupInUpstream, len(u.Spec.ServerGroups))
			for i, g := range u.Spec.ServerGroups {
				serverGroups[i] = vproxy_config.ServerGroupInUpstream{
					Name:        makeName(ns, g.Name),
					Weight:      g.Weight,
					Annotations: g.Annotations,
				}
			}
			vu := vproxy_config.Upstream{
				Base: makeBaseWithBase(ns, u.Base),
				Spec: vproxy_config.UpstreamSpec{ServerGroups: serverGroups},
			}
			for _, x := range u.Spec.ServerGroups {
				sgInUse[x.Name] = true
			}
			ret = append(ret, &vu)
		}
		for _, s := range p.sg {
			if !sgInUse[s.Metadata.Name] {
				continue
			}
			foundSg[s.Metadata.Name] = true
			vs := vproxy_config.ServerGroup{
				Base: makeBaseWithBase(ns, s.Base),
				Spec: vproxy_config.ServerGroupSpec{
					ServerGroupSelfSpec: s.Spec.ServerGroupSelfSpec,
					Servers: vproxy_config.ServerInServerGroup{
						Static: makeStaticServers(ns, s, p.ep),
					},
				},
			}
			ret = append(ret, &vs)
		}
		for _, s := range p.secg {
			if !secgInUse[s.Metadata.Name] {
				continue
			}
			foundSecg[s.Metadata.Name] = true
			vs := vproxy_config.SecurityGroup{
				Base: makeBaseWithBase(ns, s.Base),
				Spec: s.Spec,
			}
			ret = append(ret, &vs)
		}
		for _, c := range p.ck {
			if !ckInUse[c.Metadata.Name] {
				continue
			}
			crt, err := base64.StdEncoding.DecodeString(c.Data.Crt)
			if err != nil {
				Log("invalid cert format for data in Secret %v/%v", ns, c.Metadata.Name)
				continue
			}
			key, err := base64.StdEncoding.DecodeString(c.Data.Key)
			if err != nil {
				Log("invalid key format for data in Secret %v/%v", ns, c.Metadata.Name)
				continue
			}
			vc := vproxy_config.CertKey{
				Base: makeBaseWithMetadata(ns, "CertKey", c.Metadata),
				Spec: vproxy_config.CertKeySpec{Pem: vproxy_config.PemCertKey{
					Certs: []string{string(crt)},
					Key:   string(key),
				}},
			}
			ret = append(ret, &vc)
		}
		// fill in missing components
		for name := range upsInUse {
			if foundUps[name] {
				continue
			}
			vu := vproxy_config.Upstream{
				Base: makeBase(ns, "Upstream", name),
				Spec: vproxy_config.UpstreamSpec{ServerGroups: make([]vproxy_config.ServerGroupInUpstream, 0)},
			}
			ret = append(ret, &vu)
		}
		for name := range sgInUse {
			if foundSg[name] {
				continue
			}
			vs := vproxy_config.ServerGroup{
				Base: makeBase(ns, "ServerGroup", name),
				Spec: vproxy_config.ServerGroupSpec{
					ServerGroupSelfSpec: vproxy_config.ServerGroupSelfSpec{
						// these values are not important, they will not be actually used
						Timeout: 1000,
						Period:  5000,
						Up:      2,
						Down:    3,
					},
					Servers: vproxy_config.ServerInServerGroup{Static: make([]vproxy_config.StaticServer, 0)},
				},
			}
			ret = append(ret, &vs)
		}
		for name := range secgInUse {
			if foundSecg[name] {
				continue
			}
			vs := vproxy_config.SecurityGroup{
				Base: makeBase(ns, "SecurityGroup", name),
				Spec: vproxy_config.SecurityGroupSpec{
					DefaultRule: "allow",
					Rules:       make([]vproxy_config.SecurityGroupRule, 0),
				},
			}
			ret = append(ret, &vs)
		}
	}
	return ret
}

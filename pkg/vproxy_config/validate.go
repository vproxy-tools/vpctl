package vproxy_config

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"reflect"
	"strings"
)

const CurrentVersion = "vproxy/v1alpha1"

var Versions = []string{
	CurrentVersion,
}

var AvailableKind = []string{
	"TcpLb", "tcp-lb", "tl",
	"Socks5Server", "socks5-server", "socks5",
	"DnsServer", "dns-server", "dns",
	"Upstream", "upstream", "ups",
	"ServerGroup", "server-group", "sg",
	"SecurityGroup", "security-group", "secg",
	"CertKey", "cert-key", "ck",
	"SmartGroupDelegate", "smart-group-delegate", "sgd",
}

var KindMap = map[string]string{
	"server-group": "ServerGroup",
	"sg":           "ServerGroup",

	"security-group": "SecurityGroup",
	"secg":           "SecurityGroup",

	"cert-key": "CertKey",
	"ck":       "CertKey",

	"upstream": "Upstream",
	"ups":      "Upstream",

	"smart-group-delegate": "SmartGroupDelegate",
	"sgd":                  "SmartGroupDelegate",

	"tcp-lb": "TcpLb",
	"tl":     "TcpLb",

	"socks5-server": "Socks5Server",
	"socks5":        "Socks5Server",

	"dns-server": "DnsServer",
	"dns":        "DnsServer",
}

func (o *Base) Validate() error {
	available := false
	for _, ver := range Versions {
		if ver == o.ApiVersion {
			available = true
			break
		}
	}
	if !available {
		return fmt.Errorf("unsupported apiVersion: %s", o.ApiVersion)
	}
	available = false
	for _, kind := range AvailableKind {
		if kind == o.Kind {
			available = true
			break
		}
	}
	if !available {
		return fmt.Errorf("unsupported kind: %s", o.Kind)
	}
	if v, ok := KindMap[o.Kind]; ok {
		o.Kind = v
	}
	if o.Metadata.Name == "" {
		return fmt.Errorf("missing metadata.name for %+v", o)
	}
	return nil
}

func (o *Base) validateForUpdating(old *Base) error {
	return nil // nothing to check
}

func (o *TcpLb) Validate() error {
	ref := fmt.Sprintf("%s:%s", "TcpLb", o.Metadata.Name)
	if o.Spec.Address == "" {
		return fmt.Errorf("missing spec.address in %s", ref)
	}
	if o.Spec.Backend == "" {
		return fmt.Errorf("missing spec.backend in %s", ref)
	}
	if o.Spec.ListOfCertKey != nil {
		// check empty string in ListOfCertKey
		for idx, ck := range o.Spec.ListOfCertKey {
			if ck == "" {
				return fmt.Errorf("invalid element in spec.listOfCertKey[%d] in %s", idx, ref)
			}
		}
	}
	return nil
}

func (o *TcpLb) validateForUpdating(old *TcpLb) (bool, error) {
	ref := fmt.Sprintf("%s:%s", "TcpLb", o.Metadata.Name)
	if o.Spec.Address != old.Spec.Address {
		return false, fmt.Errorf("cannot modify immutable field spec.address in %s from %s to %s", ref, old.Spec.Address, o.Spec.Address)
	}
	if o.Spec.Backend != old.Spec.Backend {
		return false, fmt.Errorf("cannot modify immutable field spec.backend in %s from %s to %s", ref, old.Spec.Backend, o.Spec.Backend)
	}
	if o.Spec.Protocol != "" && o.Spec.Protocol != old.Spec.Protocol {
		return false, fmt.Errorf("cannot modify immutable field spec.protocol in %s from %s to %s", ref, old.Spec.Protocol, o.Spec.Protocol)
	}
	if (len(o.Spec.ListOfCertKey) != 0 || len(old.Spec.ListOfCertKey) != 0) &&
		!reflect.DeepEqual(o.Spec.ListOfCertKey, old.Spec.ListOfCertKey) {
		return false, fmt.Errorf("cannot modify immutable field spec.listOfCertKey in %s from %v to %v", ref, old.Spec.ListOfCertKey, o.Spec.ListOfCertKey)
	}
	update := false
	if o.Spec.SecurityGroup != "" && o.Spec.SecurityGroup != old.Spec.SecurityGroup {
		update = true
	}
	return update, nil
}

func (o *Socks5Server) Validate() error {
	ref := fmt.Sprintf("%s:%s", "Socks5Server", o.Metadata.Name)
	if o.Spec.Address == "" {
		return fmt.Errorf("missing spec.address in %s", ref)
	}
	if o.Spec.Backend == "" {
		return fmt.Errorf("missing spec.backend in %s", ref)
	}
	return nil
}

func (o *Socks5Server) validateForUpdating(old *Socks5Server) (bool, error) {
	ref := fmt.Sprintf("%s:%s", "Socks5Server", o.Metadata.Name)
	if o.Spec.Address != old.Spec.Address {
		return false, fmt.Errorf("cannot modify immutable field spec.address in %s from %s to %s", ref, old.Spec.Address, o.Spec.Address)
	}
	if o.Spec.Backend != old.Spec.Backend {
		return false, fmt.Errorf("cannot modify immutable field spec.backend in %s from %s to %s", ref, old.Spec.Backend, o.Spec.Backend)
	}
	update := false
	if o.Spec.SecurityGroup != "" && o.Spec.SecurityGroup != old.Spec.SecurityGroup {
		update = true
	}
	if o.Spec.AllowNonBackend != nil && *o.Spec.AllowNonBackend != *old.Spec.AllowNonBackend {
		update = true
	}
	return update, nil
}

func (o *DnsServer) Validate() error {
	ref := fmt.Sprintf("%s:%s", "DnsServer", o.Metadata.Name)
	if o.Spec.Address == "" {
		return fmt.Errorf("missing spec.address in %s", ref)
	}
	if o.Spec.RRSets == "" {
		return fmt.Errorf("missing spec.rrsets in %s", ref)
	}
	if o.Spec.TTL != nil && *o.Spec.TTL < 0 {
		return fmt.Errorf("invalid value for spec.ttl in %s: %d < 0", ref, o.Spec.TTL)
	}
	return nil
}

func (o *DnsServer) validateForUpdating(old *DnsServer) (bool, error) {
	ref := fmt.Sprintf("%s:%s", "DnsServer", o.Metadata.Name)
	if o.Spec.Address != old.Spec.Address {
		return false, fmt.Errorf("cannot modify immutable field spec.address in %s from %s to %s", ref, old.Spec.Address, o.Spec.Address)
	}
	if o.Spec.RRSets != old.Spec.RRSets {
		return false, fmt.Errorf("cannot modify immutable field spec.rrsets in %s from %s to %s", ref, old.Spec.RRSets, o.Spec.RRSets)
	}
	update := false
	if o.Spec.TTL != nil && *o.Spec.TTL != *old.Spec.TTL {
		update = true
	}
	if o.Spec.SecurityGroup != "" && o.Spec.SecurityGroup != old.Spec.SecurityGroup {
		update = true
	}
	return update, nil
}

func (o *StaticServer) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("missing name")
	}
	if o.Address == "" {
		return fmt.Errorf("missing address")
	}
	if o.Weight != nil && *o.Weight < 0 {
		return fmt.Errorf("invalid value for weight: %d < 0", o.Weight)
	}
	return nil
}

func (o *ServerGroup) Validate() error {
	ref := fmt.Sprintf("%s:%s", "ServerGroup", o.Metadata.Name)
	if o.Spec.Timeout <= 0 {
		return fmt.Errorf("invalid value for spec.timeout in %s: %d <= 0", ref, o.Spec.Timeout)
	}
	if o.Spec.Period <= 0 {
		return fmt.Errorf("invalid value for spec.period in %s: %d <= 0", ref, o.Spec.Period)
	}
	if o.Spec.Up <= 0 {
		return fmt.Errorf("invalid value for spec.up in %s: %d <= 0", ref, o.Spec.Up)
	}
	if o.Spec.Down <= 0 {
		return fmt.Errorf("invalid value for spec.down in %s: %d <= 0", ref, o.Spec.Down)
	}
	if o.Spec.Servers.Static != nil {
		for idx, svr := range o.Spec.Servers.Static {
			err := svr.Validate()
			if err != nil {
				return fmt.Errorf("invalid value for spec.servers.static[%d] in %s: %s", idx, ref, err.Error())
			}
		}
	}
	return nil
}

func (o *StaticServer) validateForUpdating(old *StaticServer) bool {
	if o.Weight != nil && *o.Weight != *old.Weight {
		return true
	}
	return false
}

func (o *ServerGroup) validateForUpdating(old *ServerGroup) (updateTop bool, updateSub bool, err error) {
	updateTop = false
	updateSub = false
	if o.Spec.Timeout != old.Spec.Timeout {
		updateTop = true
	}
	if o.Spec.Period != old.Spec.Period {
		updateTop = true
	}
	if o.Spec.Up != old.Spec.Up {
		updateTop = true
	}
	if o.Spec.Down != old.Spec.Down {
		updateTop = true
	}
	if o.Spec.Protocol != "" && o.Spec.Protocol != old.Spec.Protocol {
		updateTop = true
	}
	if o.Spec.Method != "" && o.Spec.Method != old.Spec.Method {
		updateTop = true
	}
	if (len(o.Metadata.Annotations) != 0 || len(old.Metadata.Annotations) != 0) &&
		!reflect.DeepEqual(o.Metadata.Annotations, old.Metadata.Annotations) {
		updateTop = true
	}
	if !reflect.DeepEqual(o.Spec.Servers, old.Spec.Servers) {
		updateSub = true
	}
	return updateTop, updateSub, nil
}

func (o *ServerGroupInUpstream) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("missing name")
	}
	if o.Weight != nil && *o.Weight < 0 {
		return fmt.Errorf("invalid value for weight: %d < 0", o.Weight)
	}
	return nil
}

func (o *ServerGroupInUpstream) validateForUpdating(old *ServerGroupInUpstream) bool {
	if o.Weight != nil && *o.Weight != *old.Weight {
		return true
	}
	if (len(o.Annotations) != 0 || len(old.Annotations) != 0) &&
		!reflect.DeepEqual(o.Annotations, old.Annotations) {
		return true
	}
	return false
}

func (o *Upstream) Validate() error {
	ref := fmt.Sprintf("%s:%s", "Upstream", o.Metadata.Name)
	for idx, sg := range o.Spec.ServerGroups {
		err := sg.Validate()
		if err != nil {
			return fmt.Errorf("invalid value for spec.serverGroups[%d] in %s: %s", idx, ref, err.Error())
		}
	}
	return nil
}

func (o *Upstream) validateForUpdating(old *Upstream) (updateTop bool, updateSub bool, err error) {
	update := false
	if !reflect.DeepEqual(o.Spec.ServerGroups, old.Spec.ServerGroups) {
		update = true
	}
	return false, update, nil
}

func (o *SecurityGroupRule) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("missing name")
	}
	if o.ClientNetwork == "" {
		return fmt.Errorf("missing clientNetwork")
	}
	if o.Protocol == "" {
		return fmt.Errorf("missing protocol")
	}
	if o.ServerPortMin < 1 {
		return fmt.Errorf("invalid value for serverPortMin: %d < 1", o.ServerPortMin)
	}
	if o.ServerPortMax > 65535 {
		return fmt.Errorf("invalid value for serverPortMax: %d > 65535", o.ServerPortMax)
	}
	if o.ServerPortMin > o.ServerPortMax {
		return fmt.Errorf("invalid value for serverPortMin and serverPortMax: %d > %d", o.ServerPortMin, o.ServerPortMax)
	}
	if o.Rule == "" {
		return fmt.Errorf("missing rule")
	}
	return nil
}

func (o *SecurityGroupRule) validateForUpdating(old *SecurityGroupRule) bool {
	return !reflect.DeepEqual(o, old)
}

func (o *SecurityGroup) Validate() error {
	ref := fmt.Sprintf("%s:%s", "SecurityGroup", o.Metadata.Name)
	if o.Spec.DefaultRule == "" {
		return fmt.Errorf("missing spec.defaultRule in %s", ref)
	}
	for idx, rule := range o.Spec.Rules {
		err := rule.Validate()
		if err != nil {
			return fmt.Errorf("invalid value for spec.rules[%d] in %s: %s", idx, ref, err.Error())
		}
	}
	return nil
}

func (o *SecurityGroup) validateForUpdating(old *SecurityGroup) (updateTop bool, updateSub bool, error error) {
	updateTop = false
	updateSub = false
	if o.Spec.DefaultRule != old.Spec.DefaultRule {
		updateTop = true
	}
	if !reflect.DeepEqual(o.Spec.Rules, old.Spec.Rules) {
		updateSub = true
	}
	return updateTop, updateSub, nil
}

func (o *CertKey) Validate() error {
	ref := fmt.Sprintf("%s:%s", "CertKey", o.Metadata.Name)
	if o.Spec.Pem.Certs == nil || len(o.Spec.Pem.Certs) == 0 {
		return fmt.Errorf("missing spec.pem.certs in %s", ref)
	}
	for idx, cert := range o.Spec.Pem.Certs {
		if cert == "" {
			return fmt.Errorf("missing value spec.pem.certs[%d] in %s", idx, ref)
		}
		o.Spec.Pem.Certs[idx] = strings.TrimSpace(cert)
	}
	if o.Spec.Pem.Key == "" {
		return fmt.Errorf("missing spec.pem.key in %s", ref)
	}
	return nil
}

func sha1(s string) string {
	sha1 := crypto.SHA1.New()
	sha1.Write([]byte(s))
	res := sha1.Sum(nil)
	return base64.StdEncoding.EncodeToString(res)
}

func (o *CertKey) validateForUpdating(old *CertKey) (bool, error) {
	ref := fmt.Sprintf("%s:%s", "CertKey", o.Metadata.Name)
	for idx, cert := range o.Spec.Pem.Certs {
		o.Spec.Pem.Certs[idx] = strings.TrimSpace(cert)
	}
	if !reflect.DeepEqual(o.Spec.Pem.Certs, old.Spec.Pem.Certs) {
		return false, fmt.Errorf("cannot modify immutable field spec.pem.certs in %s from %s to %s", ref, old.Spec.Pem.Certs, o.Spec.Pem.Certs)
	}
	if sha1(strings.TrimSpace(o.Spec.Pem.Key)) != old.Status.KeySHA1 {
		return false, fmt.Errorf("cannot modify immutable field spec.pem.key in %s from %s to %s", ref, old.Spec.Pem.Key, o.Spec.Pem.Key)
	}
	update := false
	//noinspection GoBoolExpressions
	return update, nil
}

func (o *SmartGroupDelegate) Validate() error {
	ref := fmt.Sprintf("%s:%s", "SmartGroupDelegate", o.Metadata.Name)
	if o.Spec.Service == "" {
		return fmt.Errorf("missing spec.service in %s", ref)
	}
	if o.Spec.Zone == "" {
		return fmt.Errorf("missing spec.zone in %s", ref)
	}
	if o.Spec.HandledGroup == "" {
		return fmt.Errorf("missing spec.handledGroup in %s", ref)
	}
	return nil
}

func (o *SmartGroupDelegate) validateForUpdating(old *SmartGroupDelegate) (bool, error) {
	ref := fmt.Sprintf("%s:%s", "SmartGroupDelegate", o.Metadata.Name)
	if o.Spec.Service != old.Spec.Service {
		return false, fmt.Errorf("cannot modify immutable field spec.pem.service in %s from %s to %s", ref, old.Spec.Service, o.Spec.Service)
	}
	if o.Spec.Zone != old.Spec.Zone {
		return false, fmt.Errorf("cannot modify immutable field spec.pem.zone in %s from %s to %s", ref, old.Spec.Zone, o.Spec.Zone)
	}
	if o.Spec.HandledGroup != old.Spec.HandledGroup {
		return false, fmt.Errorf("cannot modify immutable field spec.pem.handledGroup in %s from %s to %s", ref, old.Spec.HandledGroup, o.Spec.HandledGroup)
	}
	update := false
	//noinspection GoBoolExpressions
	return update, nil
}

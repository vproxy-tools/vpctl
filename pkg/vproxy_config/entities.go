package vproxy_config

import "fmt"

type Config interface {
	GetBase() Base
	Validate() error
}

type ConfigSortingArrayForApplying []Config
type ConfigSortingArrayForDeleting []Config

type Metadata struct {
	Name        string            `json:"name" yaml:"name"`
	Annotations map[string]string `json:"annotations" yaml:"annotations"`
}

type Base struct {
	ApiVersion string   `json:"apiVersion" yaml:"apiVersion"`
	Kind       string   `json:"kind" yaml:"kind"`
	Metadata   Metadata `json:"metadata" yaml:"metadata"`
}

func (o *Base) GetBase() Base {
	return *o
}

type TcpLbSpec struct {
	Address       string   `json:"address" yaml:"address"`
	Backend       string   `json:"backend" yaml:"backend"`
	Protocol      string   `json:"protocol" yaml:"protocol"`
	ListOfCertKey []string `json:"listOfCertKey" yaml:"listOfCertKey"`
	SecurityGroup string   `json:"securityGroup" yaml:"securityGroup"`
}

type TcpLb struct {
	Base `yaml:",inline"`
	Spec TcpLbSpec `json:"spec" yaml:"spec"`
}

func (o *TcpLb) GetBase() Base {
	return o.Base
}

type Socks5ServerSpec struct {
	Address         string `json:"address" yaml:"address"`
	Backend         string `json:"backend" yaml:"backend"`
	SecurityGroup   string `json:"securityGroup" yaml:"securityGroup"`
	AllowNonBackend bool   `json:"allowNonBackend" yaml:"allowNonBackend"`
}

type Socks5Server struct {
	Base `yaml:",inline"`
	Spec Socks5ServerSpec `json:"spec" yaml:"spec"`
}

func (o *Socks5Server) GetBase() Base {
	return o.Base
}

type DnsServerSpec struct {
	Address       string `json:"address" yaml:"address"`
	RRSets        string `json:"rrsets" yaml:"rrsets"`
	TTL           int    `json:"ttl" yaml:"ttl"`
	SecurityGroup string `json:"securityGroup" yaml:"securityGroup"`
}

type DnsServer struct {
	Base `yaml:",inline"`
	Spec DnsServerSpec `json:"spec" yaml:"spec"`
}

func (o *DnsServer) GetBase() Base {
	return o.Base
}

type StaticServer struct {
	Name    string `json:"name" yaml:"name"`
	Address string `json:"address" yaml:"address"`
	Weight  int    `json:"weight" yaml:"weight"`
}

type ServerInServerGroup struct {
	Static []StaticServer `json:"static" yaml:"static"`
}

type ServerGroupSpec struct {
	Timeout  int                 `json:"timeout" yaml:"timeout"`
	Period   int                 `json:"period" yaml:"period"`
	Up       int                 `json:"up" yaml:"up"`
	Down     int                 `json:"down" yaml:"down"`
	Protocol string              `json:"protocol" yaml:"protocol"`
	Method   string              `json:"method" yaml:"method"`
	Servers  ServerInServerGroup `json:"servers" yaml:"servers"`
}

type ServerStatus struct {
	Name      string `json:"name" yaml:"name"`
	Address   string `json:"address" yaml:"address"`
	Weight    int    `json:"weight" yaml:"weight"`
	CurrentIp string `json:"currentIp" yaml:"currentIp"`
	Status    string `json:"status" yaml:"status"`
}

type ServerGroupStatus struct {
	Servers []ServerStatus `json:"servers" yaml:"servers"`
}

type ServerGroup struct {
	Base   `yaml:",inline"`
	Spec   ServerGroupSpec   `json:"spec" yaml:"spec"`
	Status ServerGroupStatus `json:"status" yaml:"status"`
}

func (o *ServerGroup) GetBase() Base {
	return o.Base
}

type ServerGroupInUpstream struct {
	Name        string            `json:"name" yaml:"name"`
	Weight      int               `json:"weight" yaml:"weight"`
	Annotations map[string]string `json:"annotations" yaml:"annotations"`
}

type UpstreamSpec struct {
	ServerGroups []ServerGroupInUpstream `json:"serverGroups" yaml:"serverGroups"`
}

type Upstream struct {
	Base `yaml:",inline"`
	Spec UpstreamSpec `json:"spec" yaml:"spec"`
}

func (o *Upstream) GetBase() Base {
	return o.Base
}

type SecurityGroupRule struct {
	Name          string `json:"name" yaml:"name"`
	ClientNetwork string `json:"clientNetwork" yaml:"clientNetwork"`
	Protocol      string `json:"protocol" yaml:"protocol"`
	ServerPortMin int    `json:"serverPortMin" yaml:"serverPortMin"`
	ServerPortMax int    `json:"serverPortMax" yaml:"serverPortMax"`
	Rule          string `json:"rule" yaml:"rule"`
}

type SecurityGroupSpec struct {
	DefaultRule string              `json:"defaultRule" yaml:"defaultRule"`
	Rules       []SecurityGroupRule `json:"rules" yaml:"rules"`
}

type SecurityGroup struct {
	Base `yaml:",inline"`
	Spec SecurityGroupSpec `json:"spec" yaml:"spec"`
}

func (o *SecurityGroup) GetBase() Base {
	return o.Base
}

type PemCertKey struct {
	Certs []string `json:"certs" yaml:"certs"`
	Key   string   `json:"key" yaml:"key"`
}

type CertKeySpec struct {
	Pem PemCertKey `json:"pem" yaml:"pem"`
}

type CertKeyStatus struct {
	CertFiles []string `json:"certs" yaml:"certs"`
	KeyFile   string   `json:"key" yaml:"key"`
	KeySHA1   string   `json:"keySHA1" yaml:"keySHA1"`
}

type CertKey struct {
	Base   `yaml:",inline"`
	Spec   CertKeySpec   `json:"spec" yaml:"spec"`
	Status CertKeyStatus `json:"status" yaml:"status"`
}

func (o *CertKey) GetBase() Base {
	return o.Base
}

type SmartGroupDelegateSpec struct {
	Service      string `json:"service" yaml:"service"`
	Zone         string `json:"zone" yaml:"zone"`
	HandledGroup string `json:"handledGroup" yaml:"handledGroup"`
}

type SmartGroupDelegate struct {
	Base `yaml:",inline"`
	Spec SmartGroupDelegateSpec `json:"spec" yaml:"spec"`
}

func (o *SmartGroupDelegate) GetBase() Base {
	return o.Base
}

func NewEntity(kind string) (Config, error) {
	switch kind {
	case "TcpLb":
		return &TcpLb{}, nil
	case "Socks5Server":
		return &Socks5Server{}, nil
	case "DnsServer":
		return &DnsServer{}, nil
	case "Upstream":
		return &Upstream{}, nil
	case "ServerGroup":
		return &ServerGroup{}, nil
	case "SecurityGroup":
		return &SecurityGroup{}, nil
	case "CertKey":
		return &CertKey{}, nil
	case "SmartGroupDelegate":
		return &SmartGroupDelegate{}, nil
	default:
		return nil, fmt.Errorf("unrecognized kind: %s", kind)
	}
}

var TypeOrderForApplying = []string{
	// no dep
	"ServerGroup",
	"SecurityGroup",
	"CertKey",

	// depends on server-group
	"Upstream",
	"SmartGroupDelegate",

	// depends on multi
	"TcpLb",
	"Socks5Server",
	"DnsServer",
}

func (arr ConfigSortingArrayForApplying) Len() int {
	return len(arr)
}

func idxInTypeOrder(kind string) int {
	for idx, s := range TypeOrderForApplying {
		if s == kind {
			return idx
		}
	}
	return -1
}

func (arr ConfigSortingArrayForApplying) Less(i, j int) bool {
	a := arr[i]
	b := arr[j]
	return idxInTypeOrder(a.GetBase().Kind) < idxInTypeOrder(b.GetBase().Kind)
}

func (arr ConfigSortingArrayForApplying) Swap(i, j int) {
	c := arr[i]
	arr[i] = arr[j]
	arr[j] = c
}

func (arr ConfigSortingArrayForDeleting) Len() int {
	return len(arr)
}

func (arr ConfigSortingArrayForDeleting) Less(i, j int) bool {
	a := arr[i]
	b := arr[j]
	return idxInTypeOrder(a.GetBase().Kind) > idxInTypeOrder(b.GetBase().Kind)
}

func (arr ConfigSortingArrayForDeleting) Swap(i, j int) {
	c := arr[i]
	arr[i] = arr[j]
	arr[j] = c
}

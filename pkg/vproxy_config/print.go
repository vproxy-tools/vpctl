package vproxy_config

import (
	"fmt"
	"github.com/go-logr/logr"
	yamllib "gopkg.in/yaml.v2"
	"reflect"
	"strconv"
)

func RunTodoListAndPrint(todo []*Todo) error {
	return RunTodoList(todo, nil)
}

func RunTodoList(todo []*Todo, logger *logr.Logger) error {
	for _, t := range todo {
		err := t.F(t)
		if t.Op == OpCreate {
			if err != nil {
				err := fmt.Errorf("creating %s failed: %s", t.Ref, err.Error())
				if logger != nil {
					logger.Error(err, err.Error())
				}
				return err
			} else {
				if logger == nil {
					fmt.Println(t.Ref + " created")
				} else {
					logger.Info(t.Ref + " created")
				}
			}
		} else if t.Op == OpUpdate {
			if err != nil {
				err := fmt.Errorf("updating %s failed: %s", t.Ref, err.Error())
				if logger != nil {
					logger.Error(err, err.Error())
				}
				return err
			} else {
				if logger == nil {
					fmt.Println(t.Ref + " configured")
				} else {
					logger.Info(t.Ref + " configured")
				}
			}
		} else if t.Op == OpDelete {
			if err != nil {
				err := fmt.Errorf("deleting %s failed: %s", t.Ref, err.Error())
				if logger != nil {
					logger.Error(err, err.Error())
				}
				return err
			} else {
				if logger == nil {
					fmt.Println(t.Ref + " deleted")
				} else {
					logger.Info(t.Ref + " deleted")
				}
			}
		} else {
			if err != nil {
				err := fmt.Errorf("BUG: should not happen")
				logger.Error(err, err.Error())
				return err
			} else {
				if t.Op == Op404 {
					if logger == nil {
						fmt.Println(t.Ref + " not found")
					} else {
						logger.Info(t.Ref + " not found")
					}
				} else if t.Op == OpText {
					if logger == nil {
						fmt.Println(t.Ref + " unchanged (text)")
					} else {
						logger.Info(t.Ref + " unchanged (text)")
					}
				} else {
					if logger == nil {
						fmt.Println(t.Ref + " unchanged")
					} else {
						logger.Info(t.Ref + " unchanged")
					}
				}
			}
		}
	}
	return nil
}

// list := [col][row]
func printTable(list [][]string) {
	maxWidthPerCol := make([]int, len(list))
	for idxCol, rows := range list {
		for _, cell := range rows {
			if len(cell) > maxWidthPerCol[idxCol] {
				maxWidthPerCol[idxCol] = len(cell)
			}
		}
	}
	for idxRow := 0; idxRow < len(list[0]); idxRow += 1 {
		for idxCol, rows := range list {
			cell := rows[idxRow]
			maxWidth := maxWidthPerCol[idxCol]
			spaces := maxWidth - len(cell)
			if idxCol < len(list)-1 {
				spaces += 3 // separator: 3 spaces
			} else {
				spaces = 0 // no need to add spaces if it's last col on this row
			}
			fmt.Print(cell)
			for i := 0; i < spaces; i += 1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func PrintYaml(list interface{}) {
	v := reflect.ValueOf(list)
	for i := 0; i < v.Len(); i += 1 {
		elemV := v.Index(i)
		elem := elemV.Interface()
		bytes, err := yamllib.Marshal(elem)
		if err != nil {
			panic(err)
		}
		fmt.Println("---")
		fmt.Println(string(bytes))
	}
}

func PrintOneTcpLb(tl *TcpLb, yaml bool) {
	PrintTcpLb([]*TcpLb{tl}, yaml)
}

func PrintTcpLb(list []*TcpLb, yaml bool) {
	if yaml {
		PrintYaml(list)
		return
	}

	cols := []string{"NAME", "ADDRESS", "BACKEND", "PROTOCOL", "TLS", "SECURITY-GROUP"}
	ret := make([][]string, len(cols))
	for i := 0; i < len(cols); i += 1 {
		ret[i] = make([]string, 1+len(list))
		ret[i][0] = cols[i]
	}
	for i, tl := range list {
		idx := i + 1
		ret[0][idx] = tl.Metadata.Name
		ret[1][idx] = tl.Spec.Address
		ret[2][idx] = tl.Spec.Backend
		ret[3][idx] = tl.Spec.Protocol
		if len(tl.Spec.ListOfCertKey) > 0 {
			ret[4][idx] = "TRUE"
		} else {
			ret[4][idx] = "FALSE"
		}
		ret[5][idx] = tl.Spec.SecurityGroup
	}
	printTable(ret)
}

func PrintOneSocks5Server(socks5 *Socks5Server, yaml bool) {
	PrintSocks5Server([]*Socks5Server{socks5}, yaml)
}

func PrintSocks5Server(list []*Socks5Server, yaml bool) {
	if yaml {
		PrintYaml(list)
		return
	}

	cols := []string{"NAME", "ADDRESS", "BACKEND", "SECURITY-GROUP", "ALLOW-NON-BACKEND"}
	ret := make([][]string, len(cols))
	for i := 0; i < len(cols); i += 1 {
		ret[i] = make([]string, 1+len(list))
		ret[i][0] = cols[i]
	}
	for i, socks5 := range list {
		idx := i + 1
		ret[0][idx] = socks5.Metadata.Name
		ret[1][idx] = socks5.Spec.Address
		ret[2][idx] = socks5.Spec.Backend
		ret[3][idx] = socks5.Spec.SecurityGroup
		if *socks5.Spec.AllowNonBackend {
			ret[4][idx] = "TRUE"
		} else {
			ret[4][idx] = "FALSE"
		}
	}
	printTable(ret)
}

func PrintOneDnsServer(dns *DnsServer, yaml bool) {
	PrintDnsServer([]*DnsServer{dns}, yaml)
}

func PrintDnsServer(list []*DnsServer, yaml bool) {
	if yaml {
		PrintYaml(list)
		return
	}

	cols := []string{"NAME", "ADDRESS", "RRSETS", "TTL", "SECURITY-GROUP"}
	ret := make([][]string, len(cols))
	for i := 0; i < len(cols); i += 1 {
		ret[i] = make([]string, 1+len(list))
		ret[i][0] = cols[i]
	}
	for i, dns := range list {
		idx := i + 1
		ret[0][idx] = dns.Metadata.Name
		ret[1][idx] = dns.Spec.Address
		ret[2][idx] = dns.Spec.RRSets
		ret[3][idx] = strconv.FormatInt(int64(*dns.Spec.TTL), 10)
		ret[4][idx] = dns.Spec.SecurityGroup
	}
	printTable(ret)
}

func PrintOneUpstream(ups *Upstream, yaml bool) {
	PrintUpstream([]*Upstream{ups}, yaml)
}

func PrintUpstream(list []*Upstream, yaml bool) {
	if yaml {
		PrintYaml(list)
		return
	}

	cols := []string{"NAME", "SERVER-GROUPS"}
	ret := make([][]string, len(cols))
	for i := 0; i < len(cols); i += 1 {
		ret[i] = make([]string, 1+len(list))
		ret[i][0] = cols[i]
	}
	for i, ups := range list {
		idx := i + 1
		ret[0][idx] = ups.Metadata.Name
		ret[1][idx] = strconv.FormatInt(int64(len(ups.Spec.ServerGroups)), 10)
	}
	printTable(ret)
}

func PrintOneServerGroup(sg *ServerGroup, yaml bool) {
	PrintServerGroup([]*ServerGroup{sg}, yaml)
}

func PrintServerGroup(list []*ServerGroup, yaml bool) {
	if yaml {
		PrintYaml(list)
		return
	}

	cols := []string{"NAME", "SERVERS", "UP", "DOWN"}
	ret := make([][]string, len(cols))
	for i := 0; i < len(cols); i += 1 {
		ret[i] = make([]string, 1+len(list))
		ret[i][0] = cols[i]
	}
	for i, sg := range list {
		idx := i + 1
		ret[0][idx] = sg.Metadata.Name
		ret[1][idx] = strconv.FormatInt(int64(len(sg.Spec.Servers.Static)), 10)
		var up int64 = 0
		var down int64 = 0
		for _, svr := range sg.Status.Servers {
			if svr.Status == "UP" {
				up += 1
			} else if svr.Status == "DOWN" {
				down += 1
			}
		}
		ret[2][idx] = strconv.FormatInt(up, 10)
		ret[3][idx] = strconv.FormatInt(down, 10)
	}
	printTable(ret)
}

func PrintOneSecurityGroup(secg *SecurityGroup, yaml bool) {
	PrintSecurityGroup([]*SecurityGroup{secg}, yaml)
}

func PrintSecurityGroup(list []*SecurityGroup, yaml bool) {
	if yaml {
		PrintYaml(list)
		return
	}

	cols := []string{"NAME", "DEFAULT-RULE", "RULES"}
	ret := make([][]string, len(cols))
	for i := 0; i < len(cols); i += 1 {
		ret[i] = make([]string, 1+len(list))
		ret[i][0] = cols[i]
	}
	for i, secg := range list {
		idx := i + 1
		ret[0][idx] = secg.Metadata.Name
		ret[1][idx] = secg.Spec.DefaultRule
		ret[2][idx] = strconv.FormatInt(int64(len(secg.Spec.Rules)), 10)
	}
	printTable(ret)
}

func PrintOneCertKey(ck *CertKey, yaml bool) {
	PrintCertKey([]*CertKey{ck}, yaml)
}

func PrintCertKey(list []*CertKey, yaml bool) {
	if yaml {
		PrintYaml(list)
		return
	}

	cols := []string{"NAME"}
	ret := make([][]string, len(cols))
	for i := 0; i < len(cols); i += 1 {
		ret[i] = make([]string, 1+len(list))
		ret[i][0] = cols[i]
	}
	for i, ck := range list {
		idx := i + 1
		ret[0][idx] = ck.Metadata.Name
	}
	printTable(ret)
}

func UtilGetAndPrintOne(typ string, name string, yaml bool) error {
	switch typ {
	case "TcpLb":
		fallthrough
	case "tcp-lb":
		fallthrough
	case "tl":
		one, err := GetTcpLb(name)
		if err != nil {
			return err
		}
		PrintOneTcpLb(one, yaml)
	case "Socks5Server":
		fallthrough
	case "socks5-server":
		fallthrough
	case "socks5":
		one, err := GetSocks5Server(name)
		if err != nil {
			return err
		}
		PrintOneSocks5Server(one, yaml)
	case "DnsServer":
		fallthrough
	case "dns-server":
		fallthrough
	case "dns":
		one, err := GetDnsServer(name)
		if err != nil {
			return err
		}
		PrintOneDnsServer(one, yaml)
	case "Upstream":
		fallthrough
	case "upstream":
		fallthrough
	case "ups":
		one, err := GetUpstream(name)
		if err != nil {
			return err
		}
		PrintOneUpstream(one, yaml)
	case "ServerGroup":
		fallthrough
	case "server-group":
		fallthrough
	case "sg":
		one, err := GetServerGroup(name)
		if err != nil {
			return err
		}
		PrintOneServerGroup(one, yaml)
	case "SecurityGroup":
		fallthrough
	case "security-group":
		fallthrough
	case "secg":
		one, err := GetSecurityGroup(name)
		if err != nil {
			return err
		}
		PrintOneSecurityGroup(one, yaml)
	case "CertKey":
		fallthrough
	case "cert-key":
		fallthrough
	case "ck":
		one, err := GetCertKey(name)
		if err != nil {
			return err
		}
		PrintOneCertKey(one, yaml)
	default:
		return fmt.Errorf("unknown type %s", typ)
	}
	return nil
}

func UtilGetAndPrintList(typ string, yaml bool) error {
	switch typ {
	case "TcpLb":
		fallthrough
	case "tcp-lb":
		fallthrough
	case "tl":
		list, err := ListTcpLb()
		if err != nil {
			return err
		}
		PrintTcpLb(list, yaml)
	case "Socks5Server":
		fallthrough
	case "socks5-server":
		fallthrough
	case "socks5":
		list, err := ListSocks5Server()
		if err != nil {
			return err
		}
		PrintSocks5Server(list, yaml)
	case "DnsServer":
		fallthrough
	case "dns-server":
		fallthrough
	case "dns":
		list, err := ListDnsServer()
		if err != nil {
			return err
		}
		PrintDnsServer(list, yaml)
	case "Upstream":
		fallthrough
	case "upstream":
		fallthrough
	case "ups":
		list, err := ListUpstream()
		if err != nil {
			return err
		}
		PrintUpstream(list, yaml)
	case "ServerGroup":
		fallthrough
	case "server-group":
		fallthrough
	case "sg":
		list, err := ListServerGroup()
		if err != nil {
			return err
		}
		PrintServerGroup(list, yaml)
	case "SecurityGroup":
		fallthrough
	case "security-group":
		fallthrough
	case "secg":
		list, err := ListSecurityGroup()
		if err != nil {
			return err
		}
		PrintSecurityGroup(list, yaml)
	case "CertKey":
		fallthrough
	case "cert-key":
		fallthrough
	case "ck":
		list, err := ListCertKey()
		if err != nil {
			return err
		}
		PrintCertKey(list, yaml)
	default:
		return fmt.Errorf("unknown type %s", typ)
	}
	return nil
}

func UtilWatchAndPrint(typ string) error {
	switch typ {
	case "HealthCheck":
	case "health-check":
	case "hc":
		stop := make(chan bool)
		chnl := make(chan *HealthCheckEventChannelMessage)
		defer close(chnl)
		err := WatchHealthCheck(chnl, stop)
		if err != nil {
			return err
		}
		defer func() {
			stop <- true
			close(stop)
		}()
		for {
			msg := <-chnl
			if msg == nil {
				break
			}
			if msg.err != nil {
				return msg.err
			}
			bytes, err := yamllib.Marshal(msg.evt)
			if err != nil { // should not happen
				return err
			}
			fmt.Println("---\n" + string(bytes))
		}
	default:
		return fmt.Errorf("unknown type %s", typ)
	}
	return nil
}

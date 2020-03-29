package vpctl

const HelpStr = `Usage:
	vpctl apply -f {filename}
	vpctl get {type} [{name}] [-o wide|yaml]
	vpctl delete -f {filename}
	vpctl delete {type} {name}
Environment Variables:
	+-------------------------+---------------------+
	| KEY                     | DEFAULT             |
	+-------------------------+---------------------+
	| VPCTL_WORKING_DIRECTORY | ~/vpctl             |
	| VPCTL_VPROXY_HTTP_PORT  | 18776               |
	| VPCTL_VPROXY_HTTP_HOST  | 127.0.0.1           |
	+-------------------------+---------------------+
Resource Types:
	+--------------------+----------------------+------------+
	| TYPE               | ALIAS                | SHORT      |
	+--------------------+----------------------+------------+
	| TcpLb              | tcp-lb               | tl         |
	| Socks5Server       | socks5-server        | socks5     |
	| DnsServer          | dns-server           | dns        |
	| Upstream           | upstream             | ups        |
	| ServerGroup        | server-group         | sg         |
	| SecurityGroup      | security-group       | secg       |
	| CertKey            | cert-key             | ck         |
	+--------------------+----------------------+------------+
Watchable Types:
	+--------------------+----------------------+------------+
	| TYPE               | ALIAS                | SHORT      |
	+--------------------+----------------------+------------+
	| HealthCheck        | health-check         | hc         |
	+--------------------+----------------------+------------+`

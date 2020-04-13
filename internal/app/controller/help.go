package controller

const HelpStr = `Usage: controller --labels='{"key":"value"}'
	[--k8s-api-server=...]
	[--token=...]
	[--ca=...]
	[--ignore-ns={...},{...}]
	[--watch-ns={...},{...}]
	[--launch-time={}]
	[--debug]
Default Values:
	--k8s-api-server=https://$KUBERNETES_SERVICE_HOST:$KUBERNETES_SERVICE_PORT
	--token=/var/run/secrets/kubernetes.io/serviceaccount/token
	--ca=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
	--ignore-ns=kube-system
	--watch-ns=all
	--launch-time=5
Note:
	--k8s-api-server= can take http:// or https:// urls
	--token= and --ca= can take empty string, which means no token or ignore ca validation
	When using http protocol, --ca= is automatically ignored
	The server certificate auth is disabled if using https protocol and --ca= is empty string
Environment Variables:
	+-------------------------+---------------------+
	| KEY                     | DEFAULT             |
	+-------------------------+---------------------+
	| VPCTL_WORKING_DIRECTORY | ~/vpctl             |
	| VPCTL_VPROXY_HTTP_PORT  | 18776               |
	| VPCTL_VPROXY_HTTP_HOST  | 127.0.0.1           |
	+-------------------------+---------------------+`

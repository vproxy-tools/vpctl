GOPATH=`pwd` go-swagger generate client -f ../vproxy/doc/api.yaml --api-package=vproxy.client.api --model-package=vproxy.client.model --client-package=vproxy.client --target=./src

# need to modify some generated code
sed -i -e 's/AllowNonBackend bool `json:"allowNonBackend,omitempty"`/AllowNonBackend *bool `json:"allowNonBackend,omitempty"`/g' "src/vproxy_client_model/socks5_server_update.go"
rm src/vproxy_client_model/socks5_server_update.go-e

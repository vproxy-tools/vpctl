#!/bin/bash

set -e

mkdir -p swagger

go-swagger generate client -f https://raw.githubusercontent.com/wkgcass/vproxy/dev/doc/api.yaml --api-package=vproxy.client.api --model-package=vproxy.client.model --client-package=vproxy.client --target=./swagger

# need to modify some generated code
sed -i -e 's/AllowNonBackend bool `json:"allowNonBackend,omitempty"`/AllowNonBackend *bool `json:"allowNonBackend,omitempty"`/g' "swagger/vproxy_client_model/socks5_server_create.go"
rm -f swagger/vproxy_client_model/socks5_server_create.go-e

sed -i -e 's/AllowNonBackend bool `json:"allowNonBackend,omitempty"`/AllowNonBackend *bool `json:"allowNonBackend,omitempty"`/g' "swagger/vproxy_client_model/socks5_server_update.go"
rm -f swagger/vproxy_client_model/socks5_server_update.go-e

sed -i -e 's/Weight int64 `json:"weight,omitempty"`/Weight *int64 `json:"weight,omitempty"`/g' "swagger/vproxy_client_model/server_update.go"
rm -f swagger/vproxy_client_model/server_update.go-e

sed -i -e 's/Weight int64 `json:"weight,omitempty"`/Weight *int64 `json:"weight,omitempty"`/g' "swagger/vproxy_client_model/server_create.go"
rm -f swagger/vproxy_client_model/server_create.go-e

sed -i -e 's/Weight int64 `json:"weight,omitempty"`/Weight *int64 `json:"weight,omitempty"`/g' "swagger/vproxy_client_model/server_group_in_upstream_create.go"
rm -f swagger/vproxy_client_model/server_group_in_upstream_create.go-e

sed -i -e 's/Weight int64 `json:"weight,omitempty"`/Weight *int64 `json:"weight,omitempty"`/g' "swagger/vproxy_client_model/server_group_in_upstream_update.go"
rm -f swagger/vproxy_client_model/server_group_in_upstream_update.go-e

sed -i -e 's/TTL int64 `json:"ttl,omitempty"`/TTL *int64 `json:"ttl,omitempty"`/g' "swagger/vproxy_client_model/dns_server_create.go"
rm -f swagger/vproxy_client_model/dns_server_create.go-e

sed -i -e 's/TTL int64 `json:"ttl,omitempty"`/TTL *int64 `json:"ttl,omitempty"`/g' "swagger/vproxy_client_model/dns_server_update.go"
rm -f swagger/vproxy_client_model/dns_server_update.go-e

GOPATH=`pwd` go-swagger generate client -f ../vproxy/doc/api.yaml --api-package=vproxy.client.api --model-package=vproxy.client.model --client-package=vproxy.client --target=./src

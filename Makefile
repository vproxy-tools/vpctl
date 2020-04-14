.PHONY: all main vpctl controller test vpctl_test clean docker_dev_vpctl docker_dev_vproxy

main: vpctl controller vpctl-linux controller-linux
all: main test
vpctl:
	GOPATH=`pwd` go build -o vpctl cmd/vpctl/main.go
test: vpctl_test
vpctl_test: vpctl
	GOPATH=`pwd` go build -o vpctl_test test/vpctl/main.go
controller:
	GOPATH=`pwd` go build -o controller cmd/controller/main.go
clean:
	rm -f vpctl
	rm -f vpctl-linux
	rm -f vpctl_test
	rm -f controller
	rm -f controller-linux
	rm -f misc/dockerfiles/vpctl/vpctl
	rm -f misc/dockerfiles/vpctl/controller

docker_dev_vpctl:
	docker rmi -f wkgcass/vpctl:latest
	env GOOS=linux GOARCH=amd64 GOPATH=`pwd` go build -o vpctl-linux cmd/vpctl/main.go
	cp ./vpctl-linux misc/dockerfiles/vpctl/vpctl
	env GOOS=linux GOARCH=amd64 GOPATH=`pwd` go build -o controller-linux cmd/controller/main.go
	cp ./controller-linux misc/dockerfiles/vpctl/controller
	docker build --no-cache -t wkgcass/vpctl:latest ./misc/dockerfiles/vpctl

docker_dev_vproxy:
	docker rmi -f wkgcass/vproxy:latest
	cp ../vproxy/build/libs/vproxy.jar ./misc/dockerfiles/vproxy/vproxy.jar
	docker build --no-cache -t wkgcass/vproxy:latest ./misc/dockerfiles/vproxy

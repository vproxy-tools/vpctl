SHELL := /bin/bash
.DEFAULT: jar

OS := $(shell uname)

.PHONY: main
main: vpctl controller
.PHONY: all
all: clean main test docker-vpctl docker-vproxy
.PHONY: vpctl
vpctl:
	GOPATH=`pwd` go build -o vpctl cmd/vpctl/main.go
.PHONY: test
test: vpctl_test
.PHONY: vpctl_test
vpctl_test: vpctl
	GOPATH=`pwd` go build -o vpctl_test test/vpctl/main.go
.PHONY: controller
controller:
	GOPATH=`pwd` go build -o controller cmd/controller/main.go
.PHONY: clean
clean:
	rm -f vpctl
	rm -f vpctl-linux
	rm -f vpctl_test
	rm -f controller
	rm -f controller-linux
	rm -f misc/dockerfiles/vpctl/vpctl
	rm -f misc/dockerfiles/vpctl/controller

.PHONY: docker-vproxy-base
docker-vproxy-base:
	docker rmi -f vproxy-base:latest
	docker build --no-cache -t vproxy-base:latest ./misc/dockerfiles/vproxy-base

.PHONY: docker-vpctl
docker-vpctl:
	docker rmi -f wkgcass/vpctl:latest
	env GOOS=linux GOARCH=amd64 GOPATH=`pwd` go build -o vpctl-linux cmd/vpctl/main.go
	cp ./vpctl-linux misc/dockerfiles/vpctl/vpctl
	env GOOS=linux GOARCH=amd64 GOPATH=`pwd` go build -o controller-linux cmd/controller/main.go
	cp ./controller-linux misc/dockerfiles/vpctl/controller
	docker build --no-cache -t wkgcass/vpctl:latest ./misc/dockerfiles/vpctl

.PHONY: docker-vproxy
docker-vproxy:
ifeq ($(OS),Linux)
	cd ../vproxy && make jlink
	rm -rf misc/dockerfiles/vproxy/vproxy-runtime-linux/
	cp -r ../vproxy/build/image misc/dockerfiles/vproxy/vproxy-runtime-linux
endif
	docker rmi -f wkgcass/vproxy:latest
	cp ../vproxy/build/libs/vproxy.jar ./misc/dockerfiles/vproxy/vproxy.jar
	docker build --no-cache -t wkgcass/vproxy:latest ./misc/dockerfiles/vproxy

.PHONY: bundle
bundle: vproxy-runtime-linux

.PHONY: vproxy-runtime-linux
vproxy-runtime-linux:
	wget https://github.com/wkgcass/vproxy/releases/download/1.0.0-BETA-5/vproxy-runtime-linux.tar.gz
	tar zxf vproxy-runtime-linux.tar.gz
	rm -f vproxy-runtime-linux.tar.gz
	rm -rf misc/dockerfiles/vproxy/vproxy-runtime-linux/
	mv vproxy-runtime-linux misc/dockerfiles/vproxy/

SHELL := /bin/bash
.DEFAULT: main

OS := $(shell uname)

.PHONY: main
main: vpctl controller
.PHONY: all
all: clean main test docker-vpctl docker-vproxy
.PHONY: vpctl
vpctl:
	GO111MODULE=off GOPATH=`pwd` go build -o vpctl cmd/vpctl/main.go
.PHONY: test
test: vpctl_test
.PHONY: vpctl_test
vpctl_test: vpctl
	GO111MODULE=off GOPATH=`pwd` go build -o vpctl_test test/vpctl/main.go
.PHONY: controller
controller:
	GO111MODULE=off GOPATH=`pwd` go build -o controller cmd/controller/main.go
.PHONY: clean
clean:
	rm -f vpctl
	rm -f vpctl-linux
	rm -f vpctl_test
	rm -f controller
	rm -f controller-linux
	rm -f misc/dockerfiles/vpctl/vpctl
	rm -f misc/dockerfiles/vpctl/controller

.PHONY: docker-vproxy-runtime
docker-vproxy-runtime:
	docker rmi -f vproxyio/runtime:latest
	docker build --no-cache -t vproxyio/runtime:latest ./misc/dockerfiles/vproxy-runtime

.PHONY: docker-vproxy-base
docker-vproxy-base:
	docker rmi -f vproxyio/base:latest
	docker build --no-cache -t vproxyio/base:latest ./misc/dockerfiles/vproxy-base

.PHONY: docker-vproxy-compile
docker-vproxy-compile:
	docker rmi -f vproxyio/compile:latest
	docker build --no-cache -t vproxyio/compile:latest ./misc/dockerfiles/vproxy-compile

.PHONY: docker-vproxy-test
docker-vproxy-test:
	docker rmi -f vproxyio/test:latest
	docker build --no-cache -t vproxyio/test:latest ./misc/dockerfiles/vproxy-test

.PHONY: docker-vpctl
docker-vpctl:
	docker rmi -f vproxyio/vpctl:latest
	env GO111MODULE=off GOOS=linux GOARCH=amd64 GOPATH=`pwd` go build -o vpctl-linux cmd/vpctl/main.go
	cp ./vpctl-linux misc/dockerfiles/vpctl/vpctl
	env GO111MODULE=off GOOS=linux GOARCH=amd64 GOPATH=`pwd` go build -o controller-linux cmd/controller/main.go
	cp ./controller-linux misc/dockerfiles/vpctl/controller
	docker build --no-cache -t vproxyio/vpctl:latest ./misc/dockerfiles/vpctl

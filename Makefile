.PHONY: all main vpctl test vpctl_test clean

main: vpctl controller
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
	rm -f vpctl_test
	rm -f controller

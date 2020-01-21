.PHONY: all main vpctl test vpctl_test clean

main: vpctl
all: main test
vpctl:
	GOPATH=`pwd` go build -o vpctl cmd/vpctl/main.go
test: vpctl_test
vpctl_test: vpctl
	GOPATH=`pwd` go build -o vpctl_test test/vpctl/main.go
clean:
	rm -f vpctl
	rm -f vpctl_test

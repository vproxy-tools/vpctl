package main

import (
	"github.com/vproxy-tools/vpctl/internal/app/vpctl"
	"os"
)

func main() {
	args := os.Args[1:]
	exitCode := vpctl.Main(args)
	os.Exit(exitCode)
}

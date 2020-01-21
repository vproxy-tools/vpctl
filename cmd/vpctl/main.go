package main

import (
	"../../internal/app/vpctl"
	"os"
)

func main() {
	args := os.Args[1:]
	exitCode := vpctl.Main(args)
	os.Exit(exitCode)
}

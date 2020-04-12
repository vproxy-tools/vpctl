package main

import (
	"../../internal/app/controller"
	"os"
)

func main() {
	args := os.Args[1:]
	exitCode := controller.Main(args)
	os.Exit(exitCode)
}

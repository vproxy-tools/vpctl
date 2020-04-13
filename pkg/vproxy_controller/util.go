package vproxy_controller

import (
	"../vproxy_config"
	"encoding/json"
	"fmt"
	"os"
)

var DebugEnabled = false

func Log(f string, args ...interface{}) {
	for i, a := range args {
		if _, ok := a.(error); ok {
			continue
		}
		if _, ok := a.(int); ok {
			continue
		}
		if _, ok := a.(string); ok {
			continue
		}
		b, e := json.Marshal(a)
		if e != nil {
			continue
		}
		args[i] = string(b)
	}
	fmt.Printf(f+"\n", args...)
}

func Debug(f string, args ...interface{}) {
	if !DebugEnabled {
		return
	}
	for i, a := range args {
		if _, ok := a.(error); ok {
			continue
		}
		if _, ok := a.(int); ok {
			continue
		}
		if _, ok := a.(string); ok {
			continue
		}
		b, e := json.Marshal(a)
		if e != nil {
			continue
		}
		args[i] = string(b)
	}
	_, _ = fmt.Fprintf(os.Stderr, f+"\n", args...)
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func int2ptr(i int) *int {
	return &i
}

func merge(a, b []*vproxy_config.Todo) []*vproxy_config.Todo {
	x := make([]*vproxy_config.Todo, len(a)+len(b))
	for i, t := range a {
		x[i] = t
	}
	off := len(a)
	for i, t := range b {
		x[off+i] = t
	}
	return x
}

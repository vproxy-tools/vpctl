package vproxy_config

import (
	"os"
	"path"
	"strconv"
	"strings"
)

func GetWorkingDir() string {
	p, b := os.LookupEnv("VPCTL_WORKING_DIRECTORY")
	if !b {
		home, _ := os.UserHomeDir()
		p = path.Join(home, "vpctl")
	}
	_ = os.MkdirAll(p, os.FileMode(0755))
	return p
}

func GetHttpPort() int {
	p, b := os.LookupEnv("VPCTL_VPROXY_HTTP_PORT")
	if !b {
		return 18776
	}
	i, _ := strconv.ParseInt(p, 10, 32)
	return int(i)
}

func GetHost() string {
	h, b := os.LookupEnv("VPCTL_VPROXY_HTTP_HOST")
	if !b || h == "" {
		return "127.0.0.1"
	}
	return h
}

func CertName(s string, idx int) string {
	return path.Join(GetWorkingDir(), s+"."+strconv.FormatInt(int64(idx), 10)+".crt")
}

func KeyName(s string) string {
	return path.Join(GetWorkingDir(), s+".key")
}

func IsCertKeyName(f string, s string) bool {
	return strings.HasPrefix(f, s)
}

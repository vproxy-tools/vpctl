#!/bin/bash

set -e

cat ../config/crd/bases/app.vproxy.io_tcplbs.yaml > crd.yaml
cat ../config/crd/bases/app.vproxy.io_socks5servers.yaml >> crd.yaml
cat ../config/crd/bases/app.vproxy.io_dnsservers.yaml >> crd.yaml
cat ../config/crd/bases/app.vproxy.io_upstreams.yaml >> crd.yaml
cat ../config/crd/bases/app.vproxy.io_servergroups.yaml >> crd.yaml
cat ../config/crd/bases/app.vproxy.io_securitygroups.yaml >> crd.yaml
cat ../config/crd/bases/app.vproxy.io_certkeys.yaml >> crd.yaml

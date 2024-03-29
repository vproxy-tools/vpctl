# vpctl

This is a tool for controlling the local vproxy instance both on command line and in a Kubernetes cluster.

## compile

You need at least Go 1.19 to compile the project.

```bash
make clean
make vpctl        # command line
make docker-build # k8s controller
```

## usage

<details><summary><code>vpctl</code> command line tool</summary>

```
vpctl apply -f {filename}
vpctl get {type} [{name}] [-o wide|yaml]
vpctl delete -f {filename}
vpctl delete {type} {name}
```

</details>

<details><summary><code>controller</code> for k8s</summary>

<br>

Apply the following yamls:

1. the CRDs of vproxy
2. launch a vproxy instance and controller in namespace `vproxy-system`

```
kubectl apply -f https://raw.githubusercontent.com/vproxy-tools/vpctl/master/misc/crd.yaml
kubectl apply -f https://raw.githubusercontent.com/vproxy-tools/vpctl/master/misc/k8s-vproxy.yaml
```

and you may apply the example resources, which will start several services, and expose them via `TcpLb`, `Socks5Server` and `DnsServer`

```
kubectl apply -f https://raw.githubusercontent.com/vproxy-tools/vpctl/master/misc/cr-example.yaml
```

After the configuration, you may use various ways to access the vproxy gateway:

```shell
# exec into the vpctl container:
kubectl -n vproxy-system exec -it --container=controller `kubectl -n vproxy-system get pod | grep vproxy-gateway | grep Running | awk '{print $1}'` /bin/bash
# inside the container, you can run these commands multiple times:
/vpctl get TcpLb

# you can attach a debug container into the vproxy-gateway pod:
# kubectl debug -n vproxy-system -it --image="$IMAGE" `kubectl -n vproxy-system get pod | grep vproxy-gateway | grep Running | awk '{print $1}'` -- sh
curl 127.0.0.1
curl -H 'Host: example.com' 127.0.0.1
curl -H 'Host: example2.com' 127.0.0.1
curl -k https://127.0.0.1
curl -k -H 'Host: example.com' https://127.0.0.1
curl -k -H 'Host: example2.com' https://127.0.0.1
dig @127.0.0.1 example.com
dig @127.0.0.1 example2.com

# OR
# expose the pod using nodePort
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: vproxy-gateway-service
  namespace: vproxy-system
spec:
  selector:
    app: vproxy-gateway
  ports:
  - port: 80
    name: http
    protocol: TCP
    nodePort: 30080
  - port: 443
    name: https
    protocol: TCP
    nodePort: 30443
  - port: 1080
    name: socks5
    protocol: TCP
    nodePort: 31080
  - port: 53
    name: dns
    protocol: UDP
    nodePort: 30053
  type: NodePort
EOF
# then you can access these services, e.g.
curl 127.0.0.1:30080
```

Note: You will not be able to access the `Socks5Server manage-socks5` because your ip is not listed in the whitelist of the `SecurityGroup work-area-only`, modify the [cr-example.yaml#94](https://github.com/vproxy-tools/vpctl/blob/master/misc/cr-example.yaml#L94) to allow your ip and re-apply it, then you can use `curl --socks5 'socks5h://xxxxxx' http://example.com` to access your services.

</details>

## config

There are 7 kinds of resources in the yaml configurations:

* TcpLb: binds an address and dispatches requests
* Socks5Server: runs a standard socks5 protocol for you to manage internal services
* DnsServer: lets you customize A|AAAA records, usually used as a component of sidecar
* ServerGroup: manages a list of addresses and other metadata of servers
* Upstream: a collection of ServerGroup resources with metadata
* SecurityGroup: white or black list of clientIp+serverPort tuples
* CertKey: holds certificates chain and the private key for TLS

You may refer to the yaml configurations in directory `test/samples/`, e.g. [example.yaml](https://github.com/vproxy-tools/vpctl/blob/master/test/samples/example.yaml).

There are several use cases provided in this README doc for you to refer to.

## kubernetes

We defined a few CRDs (Custom Resource Definitions) to integrate vproxy into Kubernetes. All of them share the same definitions with vpctl yaml configurations, except that:

1. The pods may change at any time, so the static address configuration for `ServerGroup` would be useless. As a result, we let users to set k8s `Endpoints` in the `ServerGroup` resource, you may refer to the [cr-example.yaml#L54](https://github.com/vproxy-tools/vpctl/blob/master/misc/cr-example.yaml#L54). Addresses are dynamically updated according to the `Endpoints` resources. You can create a `Service` resource with the `selectors` field to let k8s automatically find pods based on their labels.

When using the vproxy gateway, you must put the gateway pods inside namespace `vproxy-system` and add labels `vproxy-app: controller` to it. Also the env variable `VPROXY_CONTROLLER_NAME` of the controller container must be set to the pod name. Finalizers will be added and removed based on the env variable and existing controller pods.

## examples

<details><summary>basic tcp</summary>

Case: You have two app servers, 10.0.10.1, 10.0.10.2. Proxy incoming connections to these two servers.

```yaml

---
apiVersion: vproxy.io/v1alpha1
kind: TcpLb
metadata:
  name: my-lb
spec:
  address: 0.0.0.0:80
  backend: ups001
  protocol: tcp

---
apiVersion: vproxy.io/v1alpha1
kind: Upstream
metadata:
  name: ups001
spec:
  serverGroups:
    - name: sg001
      weight: 10

---
apiVersion: vproxy.io/v1alpha1
kind: ServerGroup
metadata:
  name: sg001
spec:
  timeout: 1000
  period: 5000
  up: 2
  down: 3
  protocol: tcp
  method: wrr
  servers:
    static:
      - name: svr1
        address: 10.0.10.1:80
        weight: 10
      - name: svr2
        address: 10.0.10.2:80
        weight: 10

```

</details>

<details><summary>basic http</summary>

Case: You have two app servers, 10.0.10.1, 10.0.10.2, both running stateless http apps. Proxy incoming http requests to these two servers, and requests in the same connection should be sent to multiple servers.

> `protocol: http` is used for both http/2 and http/1.x, you may use `h2` or `http/1.x` as well.

```yaml

---
apiVersion: vproxy.io/v1alpha1
kind: TcpLb
metadata:
  name: my-lb
spec:
  address: 0.0.0.0:80
  backend: ups001
  protocol: http

---
apiVersion: vproxy.io/v1alpha1
kind: Upstream
metadata:
  name: ups001
spec:
  serverGroups:
    - name: sg001
      weight: 10

---
apiVersion: vproxy.io/v1alpha1
kind: ServerGroup
metadata:
  name: sg001
spec:
  timeout: 1000
  period: 5000
  up: 2
  down: 3
  protocol: tcp
  method: wrr
  servers:
    static:
      - name: svr1
        address: 10.0.10.1:80
        weight: 10
      - name: svr2
        address: 10.0.10.2:80
        weight: 10

```

</details>

<details><summary>tls offload</summary>

Case: You have two servers (10.0.10.1:80, 10.0.10.2:80) receiving http flow, and want to expose them to the internet using https.

> Choosing protocol `http`, will give `alpn: h2,http/1.1`, if you need `alpn: h2` only, set protocol to `h2`, or if you need `alpn: http/1.1` only, set to `http/1.x`

```yaml

---
apiVersion: vproxy.io/v1alpha1
kind: TcpLb
metadata:
  name: my-lb
spec:
  address: 0.0.0.0:443
  backend: ups001
  protocol: http
  listOfCertKey:
    - cert-xxx.com

---
apiVersion: vproxy.io/v1alpha1
kind: Upstream
metadata:
  name: ups001
spec:
  serverGroups:
    - name: sg001
      weight: 10

---
apiVersion: vproxy.io/v1alpha1
kind: ServerGroup
metadata:
  name: sg001
spec:
  timeout: 1000
  period: 5000
  up: 2
  down: 3
  protocol: tcp
  method: wrr
  servers:
    static:
      - name: svr1
        address: 10.0.10.1:80
        weight: 10
      - name: svr2
        address: 10.0.10.2:80
        weight: 10

---
apiVersion: vproxy.io/v1alpha1
kind: CertKey
metadata:
  name: cert-xxx.com
spec:
  pem:
    certs:
      - |
        -----BEGIN CERTIFICATE-----
        MIIDszCCApugAwIBAgIJAIvTzI2C9kiOMA0GCSqGSIb3DQEBCwUAMGIxCzAJBgNV
        BAYTAkNOMQ8wDQYDVQQIDAZ2cHJveHkxDzANBgNVBAcMBnZwcm94eTEPMA0GA1UE
        CgwGdnByb3h5MQ8wDQYDVQQLDAZ2cHJveHkxDzANBgNVBAMMBnZwcm94eTAeFw0y
        MDAxMjIwNjIyNDZaFw0yMTAxMjEwNjIyNDZaMGsxCzAJBgNVBAYTAkNOMRAwDgYD
        VQQIDAdleGFtcGxlMRAwDgYDVQQHDAdleGFtcGxlMRAwDgYDVQQKDAdleGFtcGxl
        MRAwDgYDVQQLDAdleGFtcGxlMRQwEgYDVQQDDAtleGFtcGxlLmNvbTCCASIwDQYJ
        KoZIhvcNAQEBBQADggEPADCCAQoCggEBAKy68AEc5T73gTX2acGIATG/8/sIpRyu
        1tsWTuB7R32A/qUYaY8NaF6ChrBiKF+eZmMHumkrgj68Haw508YWXj+QwGjtIYay
        iwcG4yNe2ojG+DRhhrAGX8GrNtJLBIgT+Shy6PSxjGV9D9sGGarHIcY87nPUC5Xf
        WMguAw/22/189igmNkAkSSJDASNRKjUuMz46nNsXsLTA+Fs8uFLa/uxCDOSQ7bRY
        4TMnnMFIf70xJoz4O0FyVossgHc1gTwskRS3CGX8MOsDnKrZ1zM5AB5rAs7FXNAO
        j+aO1k6SDWKBLPdQkKugiqh13idUbpa1v9lc7+HLqxzI+u27E2HoODcCAwEAAaNj
        MGEwCQYDVR0TBAIwADALBgNVHQ8EBAMCBeAwRwYDVR0RBEAwPoINKi5leGFtcGxl
        LmNvbYILZXhhbXBsZS5jb22CECouZXhhbXBsZS5jb20uaGuCDmV4YW1wbGUuY29t
        LmhrMA0GCSqGSIb3DQEBCwUAA4IBAQCeWjY5eQCVYjhxUUP75vWnS0ZciRPj+0u8
        HR/UcIlFL0FNNfi1EGN2z0wA+mqHe6nkRfl5FyD/NaUaVspsDOvpCXOtwHNR0izE
        TYHJrLphVtciEC9Ko9nHjN7O/VAAumNvnWt/UJbjZPq3q47FkAhQHRxdiSKstda+
        CU+fK5sT5CKguRK7j7un6u4vZ/cgIBIhrpp7pAhBa+JVW/8Wih7g0K364i4aaU8N
        QgVAU1Pziix3Cbejifu5zl9m9NCImDPV+mJZpZePiH4mrqGIymoxo4tqufl74im1
        RKyce75jojmg+6W0WmQAVrg3WPFNktHE0HSYhT+WLEzlHIHR9yB6
        -----END CERTIFICATE-----
      - |
        -----BEGIN CERTIFICATE-----
        MIIDwjCCAqqgAwIBAgIJANVe3FnIsItZMA0GCSqGSIb3DQEBCwUAMGIxCzAJBgNV
        BAYTAkNOMQ8wDQYDVQQIDAZ2cHJveHkxDzANBgNVBAcMBnZwcm94eTEPMA0GA1UE
        CgwGdnByb3h5MQ8wDQYDVQQLDAZ2cHJveHkxDzANBgNVBAMMBnZwcm94eTAgFw0y
        MDAxMjIwNjIwMDRaGA8yMTE5MTIyOTA2MjAwNFowYjELMAkGA1UEBhMCQ04xDzAN
        BgNVBAgMBnZwcm94eTEPMA0GA1UEBwwGdnByb3h5MQ8wDQYDVQQKDAZ2cHJveHkx
        DzANBgNVBAsMBnZwcm94eTEPMA0GA1UEAwwGdnByb3h5MIIBIjANBgkqhkiG9w0B
        AQEFAAOCAQ8AMIIBCgKCAQEAvxOewUhOqAzb/lRnbjQRBYgohvizlOJg5Julty9o
        /RdHe2qTw0EBouvozN2nYPl8awJKofT/N0UWA2ST0DasS4bS6c/h0dnz+14rFWMN
        ruBSlJvivCrQJZz4y3oMBkixuuSZibQib28mDGfPnOu50H/wsKhBzdNtVGQxDsSP
        fl0Xl4TnxK+sN7IG5as5cg7Gm3J8HJfO3AXyY0jVhKHbFRkUKMVqy3v3aYFtxmro
        ikNaeDv2qhhlmojYA1isRiK/+m2n3SoADQXaeGD0SMvPlblQ7x0EBlksfmQD7/Lu
        MqqItUr9mdIm8acsp9xFyylb66uTjARTVh5eZ7TD2v0XSQIDAQABo3kwdzAdBgNV
        HQ4EFgQUbku+JaaVtvRd9BDACk+FFl2dRqAwHwYDVR0jBBgwFoAUbku+JaaVtvRd
        9BDACk+FFl2dRqAwEgYDVR0TAQH/BAgwBgEB/wIBAzAOBgNVHQ8BAf8EBAMCAQYw
        EQYJYIZIAYb4QgEBBAQDAgEGMA0GCSqGSIb3DQEBCwUAA4IBAQBm9aUh3qf0A+Qj
        Q5LHetjoZOmewaUvRCAfuSnQZx2gOKr+JbOEbXuhoC+/oHxxwT4wVFL9x5Kb34Dk
        Tasn7BQUmtn8mFIQ6ryiuXKkjnzitfVOA3bSd2jvfrYHOpvn4oxvLi01deqpohhP
        LtfsF/gPEujCD5bm6u3s7i7kn5bFZC45b6yg3rcLeI9VSEm97Guza98HxaUrQA2W
        5dGbcerz4xSXaNjbFd7MHBWqy0fh/i82yWSONxPr7RBgo5Gv/usLvZQgUBy9Qd02
        eTB9efAr/JnF1SfqHAP++y35iWvY1kiWL56jSbvftrEBJdRfPhg3UP/8IkHLWi5X
        5oBm53Ci
        -----END CERTIFICATE-----
    key: |
      -----BEGIN PRIVATE KEY-----
      MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCsuvABHOU+94E1
      9mnBiAExv/P7CKUcrtbbFk7ge0d9gP6lGGmPDWhegoawYihfnmZjB7ppK4I+vB2s
      OdPGFl4/kMBo7SGGsosHBuMjXtqIxvg0YYawBl/BqzbSSwSIE/kocuj0sYxlfQ/b
      BhmqxyHGPO5z1AuV31jILgMP9tv9fPYoJjZAJEkiQwEjUSo1LjM+OpzbF7C0wPhb
      PLhS2v7sQgzkkO20WOEzJ5zBSH+9MSaM+DtBclaLLIB3NYE8LJEUtwhl/DDrA5yq
      2dczOQAeawLOxVzQDo/mjtZOkg1igSz3UJCroIqodd4nVG6Wtb/ZXO/hy6scyPrt
      uxNh6Dg3AgMBAAECggEAIFd0PtEZP3v6To9P5/O+PYFyfmt09vTpt+XqaO/MR/yS
      MXlrsmRRSXjpHTeT1auEAkDdL3n9DiAM8kV1mJ5xkTdaP4s+/siJbYxllSbahkez
      C6PDI0xeO5YCUsfHFm+e9GrnoYFpB0uuX1U6ExbsBQE9qKRocux55ClxW1jeggzH
      yrCD46OeNw7+y49JOBUoacEG/j3jQAbp8s158u/qhtz7k/uRULFl7xI143PDph1x
      zGOt2NVgOOrql5cMvGNz8DVp827OR+bONwuTBU+7CLynG7y1b2qPNtxVLTu7xKFF
      fbl7n29IkcJ/6586re45uAgGQK/S4otAD8mYa/5boQKBgQDkV/aUZ4vWK+s/JfUs
      rVwlK59bKxOD61HYlvUi/1cZS3Od3lVGKjQ5zccqCHZxKe5VCV4n9ekoiAP9iZlL
      FWYesnt7IMi9bRdmBEbNoMJrqo8zDeMmqNDiDWMNULUiE5IDkLOzfY9ld/NbIAhR
      5+wgnaKXZtZHQcjoXSDNo3iDEQKBgQDBpp9GTK7brSdhovDeGrloCEU/Pbb+1Eds
      gPglLuz9tS2ZCqHEyHGUoY+o/dcFII8HolJopwhlb+OdJDFz76dAQkkGhkvNpL7V
      9R/7I0szPvASa7+zrbKpP/fxGQxcGkb32mtZyqPDwxtljDUNk0Bgfs4jTwAUl6Io
      vZazXMj2xwKBgQCVwQ7JU4OFSbZ16sn5rBSDmDFh1EVvPhSmbJKGilmwECjaP2dD
      pggsZMWazoQHQY26HXOv13o7h8C+NdDgSj94IGwVW3HrsbEnyeQ5lZYMkIZr4E66
      GvsrVcZBhE3W9GjNh8gjDlTOIjXq7H4oYWceGOP6UYp0nzNJGVKbKvutUQKBgCvC
      1ZdzWMh31sBvq/LlIyTpSYzDC4mGuyU/99OfSRsESGufRXNMwK4P3IEZ6+9Srj/R
      ZMIVjQYvRMaMGUjTzX3t/MamrpaoNh/vpux/y0ynWmUvSED4bbllpUgsmuhtX8A+
      8ad27Y8dliFaj9qjfhbQUREVlzUQFysRvO6HdzqdAoGBAI0tlQsM7inOw6oKi/7W
      waammR8wHOXamCHq16y54ZgvibpHuR+XefvVXPoPI153fIJ7nUF9Ib5p2MSRHcAv
      FReuaQoLf3ARwOUgqMyJXFQ3Kc/6R7OzbQeagiLAsfA99ke3DxRQNMtgV2ryHHQy
      xrRx/RmwjWzCqHjobHFQN2ry
      -----END PRIVATE KEY-----


```

</details>

<details><summary>internal server maintenance</summary>

Case: You have several servers and want to access them through ssh, and only allow working area (e.g. 69.0.0.0/8) to log in into the servers. You may use vproxy Socks5Server to expose all servers on one port.

> The server-group and upstream config are emitted, see basic examples.

```yaml

---
apiVersion: vproxy.io/v1alpha1
kind: Socks5Server
metadata:
  name: socks5-001
spec:
  address: 0.0.0.0:1080
  backend: ups001
  securityGroup: work-place

---
apiVersion: vproxy.io/v1alpha1
kind: SecurityGroup
metadata:
  name: work-place
spec:
  defaultRule: deny
  rules:
    - name: city-1
      clientNetwork: 69.0.0.0/8
      protocol: TCP
      serverPortMin: 1080
      serverPortMax: 1080
      rule: allow

```

</details>

<details><summary>multiple hosts in one lb</summary>

Case: You have multiple groups of http servers and each group serving one host (e.g. www.example.com, foo.bar.com). And you want to expose them on one load balancer.

> The LB will choose one cert based on the SNI, otherwise, the first cert will be chosen.  
> Certificate configurations are omitted here, see `tls offload` for more info.

```yaml

---
apiVersion: vproxy.io/v1alpha1
kind: TcpLb
metadata:
  name: tl001
spec:
  address: 0.0.0.0:443
  backend: ups001
  protocol: http
  listOfCertKey:
    - cert-example.com
    - cert-foo.bar.com

---
apiVersion: vproxy.io/v1alpha1
kind: Upstream
metadata:
  name: ups001
spec:
  serverGroups:
    - name: sg-example
      weight: 10
    - name: sg-foo.bar
      weight: 10

---
apiVersion: vproxy.io/v1alpha1
kind: ServerGroup
metadata:
  name: sg-example
  annotations:
    host: example.com
spec:
  timeout: 1000
  period: 5000
  up: 2
  down: 3
  protocol: tcp
  method: wrr
  servers:
    static:
      - name: svr1
        address: 10.0.10.1:8080
        weight: 10
      - name: svr2
        address: 10.0.10.2:8080
        weight: 10

---
apiVersion: vproxy.io/v1alpha1
kind: ServerGroup
metadata:
  name: sg-foo.bar
  annotations:
    host: foo.bar.com
spec:
  timeout: 1000
  period: 5000
  up: 2
  down: 3
  protocol: tcp
  method: wrr
  servers:
    static:
      - name: svr1
        address: 10.0.10.1:8989
        weight: 10
      - name: svr2
        address: 10.0.10.2:8989
        weight: 10

```

</details>

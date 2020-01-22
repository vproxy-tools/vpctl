# vpctl

This is a tool for controlling the local vproxy instance.

## compile

```
GOPATH=`pwd` go build -o vpctl cmd/vpctl/main.go
```

If you are using windows, you may have to replace `pwd` with full path to the project root directory.

## usage

```
vpctl apply -f {filename}
vpctl get {type} [{name}] [-o wide|yaml]
vpctl delete -f {filename}
vpctl delete {type} {name}
```

## config

You may refer to the yaml configurations in directory `test/samples/`, e.g. [example.yaml](https://github.com/vproxy-tools/vpctl/blob/master/test/samples/example.yaml).

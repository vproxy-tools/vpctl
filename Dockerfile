FROM ubuntu:22.04
WORKDIR /
COPY bin/manager /manager
COPY vpctl-linux /vpctl

ENTRYPOINT ["/manager"]

FROM ubuntu:22.04
WORKDIR /
COPY bin/manager /manager

ENTRYPOINT ["/manager"]

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY bin/manager /manager
USER 65532:65532

ENTRYPOINT ["/manager"]

FROM ubuntu:22.04

RUN apt-get update -y
RUN apt-get install -y iproute2 iputils-ping libelf-dev wget curl redis-tools

RUN wget https://download.java.net/java/GA/jdk21.0.1/415e3f918a1f4062a0074a2794853d0d/12/GPL/openjdk-21.0.1_linux-x64_bin.tar.gz && tar zxf openjdk-21.0.1_linux-x64_bin.tar.gz && rm -f openjdk-21.0.1_linux-x64_bin.tar.gz && ./jdk-21.0.1/bin/jlink --add-modules jdk.unsupported,jdk.crypto.ec,jdk.crypto.cryptoki --output ./vproxy-runtime && rm -rf ./jdk-21.0.1

ENV JAVA_HOME="/vproxy-runtime"
ENV PATH="${PATH}:/vproxy-runtime/bin"

COPY init.sh /init.sh

RUN chmod +x /init.sh

ENTRYPOINT ["/init.sh"]

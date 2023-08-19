FROM ubuntu:22.04

COPY sources.list /etc/apt/sources.list

RUN apt-get update -y
RUN apt-get install -y iproute2 iputils-ping libelf-dev wget curl redis-tools

RUN wget https://download.java.net/java/GA/jdk21/fd2272bbf8e04c3dbaee13770090416c/35/GPL/openjdk-21_linux-x64_bin.tar.gz && tar zxf openjdk-21_linux-x64_bin.tar.gz && rm -f openjdk-21_linux-x64_bin.tar.gz && ./jdk-21/bin/jlink --add-modules jdk.unsupported,jdk.crypto.ec,jdk.crypto.cryptoki --output ./vproxy-runtime && rm -rf ./jdk-21

ENV JAVA_HOME="/vproxy-runtime"
ENV PATH="${PATH}:/vproxy-runtime/bin"

COPY init.sh /init.sh

RUN chmod +x /init.sh

ENTRYPOINT ["/init.sh"]

FROM ubuntu:22.04

COPY sources.list /etc/apt/sources.list

RUN apt-get update -y
RUN apt-get install -y iproute2 iputils-ping libelf-dev wget curl redis-tools

RUN wget https://github.com/adoptium/temurin17-binaries/releases/download/jdk-17%2B35/OpenJDK17-jdk_x64_linux_hotspot_17_35.tar.gz && tar zxf OpenJDK17-jdk_x64_linux_hotspot_17_35.tar.gz && rm -f OpenJDK17-jdk_x64_linux_hotspot_17_35.tar.gz && ./jdk-17+35/bin/jlink --add-modules jdk.unsupported,jdk.crypto.ec,jdk.crypto.cryptoki --output ./vproxy-runtime && rm -rf ./jdk-17+35

ENV JAVA_HOME="/vproxy-runtime"
ENV PATH="${PATH}:/vproxy-runtime/bin"

COPY init.sh /init.sh

RUN chmod +x /init.sh

ENTRYPOINT ["/init.sh"]

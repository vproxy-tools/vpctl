FROM ubuntu:22.04

RUN apt-get update -y
RUN apt-get install -y iproute2 iputils-ping libelf-dev wget curl redis-tools

RUN wget -O jdk.tar.gz 'https://download.java.net/java/GA/jdk22.0.1/c7ec1332f7bb44aeba2eb341ae18aca4/8/GPL/openjdk-22.0.1_linux-aarch64_bin.tar.gz' && tar zxf jdk.tar.gz && rm -f jdk.tar.gz && ./jdk-22.0.1/bin/jlink --add-modules jdk.unsupported,jdk.crypto.cryptoki --output ./vproxy-runtime && rm -rf ./jdk-22.0.1

ENV JAVA_HOME="/vproxy-runtime"
ENV PATH="${PATH}:/vproxy-runtime/bin"

COPY init.sh /init.sh

RUN chmod +x /init.sh

ENTRYPOINT ["/init.sh"]

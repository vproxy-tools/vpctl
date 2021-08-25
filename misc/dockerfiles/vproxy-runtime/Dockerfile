FROM ubuntu:20.04

COPY sources.list /etc/apt/sources.list

RUN apt-get update -y
RUN apt-get install -y iproute2 iputils-ping libelf-dev wget

RUN wget https://github.com/adoptium/temurin11-binaries/releases/download/jdk-11.0.12%2B7/OpenJDK11U-jdk_x64_linux_hotspot_11.0.12_7.tar.gz && tar zxf OpenJDK11U-jdk_x64_linux_hotspot_11.0.12_7.tar.gz && rm -f OpenJDK11U-jdk_x64_linux_hotspot_11.0.12_7.tar.gz && ./jdk-11.0.12+7/bin/jlink --add-modules jdk.unsupported,jdk.crypto.ec,jdk.crypto.cryptoki --output ./vproxy-runtime && rm -rf ./jdk-11.0.12+7

ENV JAVA_HOME="/vproxy-runtime"
ENV PATH="${PATH}:/vproxy-runtime/bin"

ENTRYPOINT ["java", "--add-exports=java.base/jdk.internal.misc=ALL-UNNAMED", "-XX:+CriticalJNINatives"]

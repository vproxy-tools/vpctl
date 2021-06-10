FROM vproxy-base:latest

RUN wget https://github.com/AdoptOpenJDK/openjdk11-binaries/releases/download/jdk-11.0.11%2B9/OpenJDK11U-jdk_x64_linux_hotspot_11.0.11_9.tar.gz && tar zxf OpenJDK11U-jdk_x64_linux_hotspot_11.0.11_9.tar.gz && rm -f OpenJDK11U-jdk_x64_linux_hotspot_11.0.11_9.tar.gz
ENV JAVA_HOME="/jdk-11.0.11+9"
RUN mkdir /vproxy
WORKDIR /vproxy

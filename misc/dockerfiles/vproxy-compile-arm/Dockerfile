FROM vproxyio/base:latest

RUN apt update
RUN apt install -y \
	libelf-dev zlib1g-dev libpcap-dev

RUN apt install -y wget apt-transport-https
RUN mkdir -p /etc/apt/keyrings
RUN wget -O - https://packages.adoptium.net/artifactory/api/gpg/key/public | tee /etc/apt/keyrings/adoptium.asc
RUN echo "deb [signed-by=/etc/apt/keyrings/adoptium.asc] https://packages.adoptium.net/artifactory/deb $(awk -F= '/^VERSION_CODENAME/{print$2}' /etc/os-release) main" | tee /etc/apt/sources.list.d/adoptium.list
RUN apt update
RUN apt install -y temurin-22-jdk

ENV JAVA_HOME="/usr/lib/jvm/temurin-22-jdk-arm64"

RUN apt-get install -y software-properties-common
RUN apt-add-repository -y ppa:lttng/stable-2.13
RUN apt-get update
RUN apt-get install -y liblttng-ust-dev lttng-tools

RUN mkdir -p /gradle/gradle/wrapper
ADD gradlew                   /gradle/gradlew
ADD gradle-wrapper.jar        /gradle/gradle/wrapper/gradle-wrapper.jar
ADD gradle-wrapper.properties /gradle/gradle/wrapper/gradle-wrapper.properties
RUN cd /gradle && chmod +x ./gradlew && ./gradlew --version

RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- --default-toolchain nightly -y --verbose
ENV PATH="$PATH:/root/.cargo/bin"

RUN mkdir /vproxy
WORKDIR /vproxy

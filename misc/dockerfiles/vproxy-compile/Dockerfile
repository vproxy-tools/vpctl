FROM vproxyio/base:latest

RUN wget https://github.com/graalvm/graalvm-ce-builds/releases/download/vm-21.2.0/graalvm-ce-java11-linux-amd64-21.2.0.tar.gz && tar zxf graalvm-ce-java11-linux-amd64-21.2.0.tar.gz && rm -f graalvm-ce-java11-linux-amd64-21.2.0.tar.gz
RUN /graalvm-ce-java11-21.2.0/bin/gu install native-image

ENV JAVA_HOME="/graalvm-ce-java11-21.2.0"
ENV PATH="${PATH}:/graalvm-ce-java11-21.2.0/bin"

RUN mkdir -p /gradle/gradle/wrapper
ADD gradlew                   /gradle/gradlew
ADD gradle-wrapper.jar        /gradle/gradle/wrapper/gradle-wrapper.jar
ADD gradle-wrapper.properties /gradle/gradle/wrapper/gradle-wrapper.properties
RUN cd /gradle && chmod +x ./gradlew && ./gradlew --version

RUN mkdir /vproxy
WORKDIR /vproxy

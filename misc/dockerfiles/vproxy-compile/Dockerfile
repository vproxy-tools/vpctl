FROM vproxyio/base:latest

RUN wget https://download.java.net/java/GA/jdk21/fd2272bbf8e04c3dbaee13770090416c/35/GPL/openjdk-21_linux-x64_bin.tar.gz && tar zxf openjdk-21_linux-x64_bin.tar.gz && rm -f openjdk-21_linux-x64_bin.tar.gz

ENV JAVA_HOME="/jdk-21"
ENV PATH="${PATH}:/jdk-21/bin"

RUN mkdir -p /gradle/gradle/wrapper
ADD gradlew                   /gradle/gradlew
ADD gradle-wrapper.jar        /gradle/gradle/wrapper/gradle-wrapper.jar
ADD gradle-wrapper.properties /gradle/gradle/wrapper/gradle-wrapper.properties
RUN cd /gradle && chmod +x ./gradlew && ./gradlew --version

RUN mkdir /vproxy
WORKDIR /vproxy

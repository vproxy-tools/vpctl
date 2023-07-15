FROM vproxyio/base:latest

RUN wget https://download.java.net/java/GA/jdk17.0.2/dfd4a8d0985749f896bed50d7138ee7f/8/GPL/openjdk-17.0.2_linux-x64_bin.tar.gz && tar zxf openjdk-17.0.2_linux-x64_bin.tar.gz && rm -f openjdk-17.0.2_linux-x64_bin.tar.gz
RUN wget https://download.java.net/java/early_access/jdk21/31/GPL/openjdk-21-ea+31_linux-x64_bin.tar.gz && tar zxf openjdk-21-ea+31_linux-x64_bin.tar.gz && rm -f openjdk-21-ea+31_linux-x64_bin.tar.gz

ENV JAVA_HOME="/jdk-17.0.2"
ENV JAVA_HOME_21="/jdk-21"
ENV PATH="${PATH}:/jdk-17.0.2/bin"

RUN mkdir -p /gradle/gradle/wrapper
ADD gradlew                   /gradle/gradlew
ADD gradle-wrapper.jar        /gradle/gradle/wrapper/gradle-wrapper.jar
ADD gradle-wrapper.properties /gradle/gradle/wrapper/gradle-wrapper.properties
RUN cd /gradle && chmod +x ./gradlew && ./gradlew --version

RUN mkdir -p /root/.gradle
ADD modify-gradle-compiler-args-agent.jar /root/.gradle/modify-gradle-compiler-args-agent.jar

RUN mkdir /vproxy
WORKDIR /vproxy

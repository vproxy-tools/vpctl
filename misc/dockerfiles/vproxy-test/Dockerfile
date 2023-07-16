FROM vproxyio/compile:latest

RUN cd /tmp && git clone https://github.com/wkgcass/vproxy --depth=1 --branch=dev && cd /tmp/vproxy && ./gradlew clean compileTestJava && cd /tmp && rm -rf /tmp/vproxy

FROM vproxyio/compile:latest

RUN cd /tmp && \
	git clone https://github.com/wkgcass/vproxy --depth=1 --branch=dev && \
	cd /tmp/vproxy && \
	git submodule update --init --recursive --recommend-shallow && \
	./gradlew clean compileTestJava && \
	make fubuki && \
	cd /tmp && rm -rf /tmp/vproxy

FROM ubuntu:22.04

RUN apt update -y
RUN DEBIAN_FRONTEND="noninteractive" apt install -y \
	lsof procps net-tools iproute2 ethtool tcpdump cpuset \
	netcat-traditional socat dnsutils iptables arping telnet \
	strace curl iputils-ping wget ca-certificates kmod \
	openssl libssl-dev \
	build-essential pkg-config autoconf \
	cmake binutils clang llvm m4 gdb lldb \
	vim tzdata uuid python3 python-is-python3 htop git
ENV TZ="Asia/Shanghai"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

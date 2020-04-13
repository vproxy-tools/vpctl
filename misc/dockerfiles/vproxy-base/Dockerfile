FROM ubuntu:16.04

RUN apt-get update -y
RUN apt-get install -y lsof procps net-tools iproute2 tcpdump strace curl iputils-ping wget build-essential autoconf netcat binutils socat dnsutils vim tzdata
ENV TZ="Asia/Shanghai"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

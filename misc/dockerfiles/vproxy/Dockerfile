FROM vproxy-base:latest

ADD vproxy-runtime-linux /runtime
ADD vproxy.jar /vproxy.jar
ENTRYPOINT ["/runtime/bin/java", "-jar", "/vproxy.jar"]

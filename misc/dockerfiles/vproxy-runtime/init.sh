#!/bin/bash
set -e

java --add-exports=java.base/jdk.internal.misc=ALL-UNNAMED -XX:+CriticalJNINatives -Djava.library.path=/usr/lib/`uname -m`-linux-gnu $@

#!/bin/bash
set -e

exec java --enable-native-access=ALL-UNNAMED --add-exports=java.base/jdk.internal.misc=ALL-UNNAMED -Djava.library.path="/usr/lib/`uname -m`-linux-gnu:/" $@

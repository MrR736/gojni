#!/usr/bin/env bash
set -e

# Find JAVA_HOME if it is not already set
if [ -z "$JAVA_HOME" ]; then
    if command -v java >/dev/null 2>&1; then
        JAVA_BIN="$(command -v java)"
        JAVA_REAL="$(readlink -f "$JAVA_BIN" 2>/dev/null || echo "$JAVA_BIN")"
        JAVA_HOME="$(dirname "$(dirname "$JAVA_REAL")")"
    else
        echo "Error: Java/JDK not found. Set JAVA_HOME."
        exit 1
    fi
fi

JNI_INCLUDE="$JAVA_HOME/include"

case "$(uname -s)" in
    Linux*)  JNI_PLATFORM="$JNI_INCLUDE/linux" ;;
    Darwin*) JNI_PLATFORM="$JNI_INCLUDE/darwin" ;;
    *)
        echo "Unsupported OS. Configure JNI paths manually."
        exit 1
        ;;
esac

if [ ! -f "$JNI_INCLUDE/jni.h" ]; then
    echo "Error: jni.h not found at $JNI_INCLUDE/jni.h"
    echo "Check that JAVA_HOME points to a JDK."
    exit 1
fi

export JAVA_HOME
export CGO_ENABLED=1
export CGO_CFLAGS="-I$JNI_INCLUDE -I$JNI_PLATFORM ${CGO_CFLAGS:-}"

echo "JAVA_HOME=$JAVA_HOME"
echo "JNI include=$JNI_INCLUDE"
echo "JNI platform=$JNI_PLATFORM"
echo "CGO_CFLAGS=$CGO_CFLAGS"
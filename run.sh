#!/bin/bash

cmd=$1

case "$cmd" in
    --build)
        go build
        ;;
    --serve)
        ./DomusPopuli-API
        ;;
    *)
        echo "Unknown command '$cmd'"
        exit 1
        ;;
esac
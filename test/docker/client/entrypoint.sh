#!/usr/bin/env sh

expect /init.sh

export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/go

cd /root/go/src/golang.org/x/crypto/ssh/gss/test/docker/client
exec go run main.go
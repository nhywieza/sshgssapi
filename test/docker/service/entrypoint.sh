#!/usr/bin/env sh

expect /genkeytab.sh

export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/go

cd /root/go/src/golang.org/x/crypto/ssh/gss/test/docker/service
exec go run main.go
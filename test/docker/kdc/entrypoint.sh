#!/usr/bin/env sh

set -e
service krb5-kdc start
service krb5-admin-server start

while true; do
    sleep 1
done

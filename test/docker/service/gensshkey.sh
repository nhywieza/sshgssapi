#!/usr/bin/expect

spawn ssh-keygen
expect "Enter file in which to save the key"
send "\r"
expect "Enter passphrase"
send "\r"
expect "Enter same passphrase again"
send "\r"
expect eof


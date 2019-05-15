#!/usr/bin/expect

spawn kinit admin/admin@EXAMPLE.COM
expect "Password for admin/admin@EXAMPLE.COM"
send "admin\r"
expect eof#!/usr/bin/env bash
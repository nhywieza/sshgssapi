#!/usr/bin/expect

spawn kinit admin/admin@EXAMPLE.COM
expect "Password for admin/admin@EXAMPLE.COM"
send "admin\r"
expect eof


spawn kadmin
expect "Password for admin/admin@EXAMPLE.COM"
send "admin\r"
expect "kadmin"
send "add_principal \"host/service@EXAMPLE.COM\"\r"
expect "Enter password for principal"
send "admin\r"
expect "Re-enter password for principal"
send "admin\r"
expect "kadmin"
send "ktadd \"host/service@EXAMPLE.COM\"\r"

expect eof


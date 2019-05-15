#!/usr/bin/expect

spawn kadmin.local
expect "kadmin.local:"
send "addprinc admin/admin@EXAMPLE.COM\r"
expect "Enter password"
send "admin\r"
expect "Re-enter"
send "admin\r"
expect "kadmin.local:"
send "exit\r"
expect eof

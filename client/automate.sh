#!/usr/bin/expect -f
 
set timeout -1
 
spawn ./client /home/bhegde/go/src/GoChat/client/pub_key /home/bhegde/go/src/GoChat/client/priv_key
 
expect "Symmetric Key exchange successful\r"
 
send -- "~~\r"
 
expect "enter '~~4' to select a user by username\r"
 
send -- "~~2\r"
 
expect "username >>\r"
 
send -- "bhegde\r"

expect "Signup was successful. You can now select a target and send messages\r"

send -- "~~4\r"

expect "target username >>\r"

send -- "potter\r"

expect "Target user is set. Target public key saved.\r"

send "hello, world\r"

expect "potter : hello\r"
 
#expect eof

#!/bin/bash
expect -c "
spawn openssl genrsa -des3 -out server.key -rand sha1.dat 2048
expect -re \"Enter.*\"
send \"1111\n\"
expect \"Verifying - Enter pass phrase for server.key:\"
send \"1111\n\"
expect
"


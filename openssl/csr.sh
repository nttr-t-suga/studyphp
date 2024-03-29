#!/bin/bash
expect -c "
spawn openssl req -new -key server.key -out server.csr
expect \"Enter pass phrase for server.key:\"
send \"1111\n\"
expect -re \"Country Name.*\"
send \"JP\n\"
expect -re \"State or Province Name.*:\"
send \"Tokyo\n\"
expect -re \"Locality Name.*:\"
send \"Chiyoda-ku\n\"
expect -re \"Organization Name.*:\"
send \"NTT COMMUNICAYTIONS CORPORATION\n\"
expect -re \"Organizational Unit Name.*:\"
send \"NetworkServices\n\"
expect -re \"Common Name.*:\"
send \"server\n\"
expect -re \"Email Address.*:\"
send \"\n\"
expect -re \"A challenge password.*:\"
send \"\n\"
expect -re \"An optional company name.*:\"
send \"\n\"
expect
"


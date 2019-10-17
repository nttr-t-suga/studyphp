# openssl genrsa -des3 -out suga.key -passout pass:suga -rand sha1.dat 2048
# openssl rsa -passin pass:suga -in suga.key -out suga.key
# openssl req -new -key suga.key -out suga.csr -subj "/C=JP/ST=Tokyo/L=Chiyoda-ku/O=NTT COMMUNICATIONS CORPORATION/OU=NetworkServices/CN=sugakunn"





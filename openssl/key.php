<?php
echo "test\n";
$name = "kokoiti" . ".key";
echo $name . "\n";

// $openssl = "openssl genrsa -des3 -out " . "server.key" . " -rand sha1.dat 2048";

exec("openssl genrsa -des3 -out " . $name . " -rand sha1.dat 2048");
?>
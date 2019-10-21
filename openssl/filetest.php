<?php
// $list = file(test);
// $list = new SplFileObject('./test.txt');

// echo $list->fgets();

// echo $list->fgets();
// $array = array();
// $array[0] = 0;
// $array[1] = 1;
// echo $array[1];

$fp = fopen('test.txt', 'r');
$csr_option = fgets($fp);

while ($name = fgetcsv($fp)) {
    echo $name[0];
    exec('openssl genrsa -des3 -out ' . $name[0] . '.key -passout pass:hoge -rand sha1.dat 2048');
    exec('openssl rsa -passin pass:hoge -in ' . $name[0] . '.key -out ' . $name[0] . '.key');
    exec('openssl req -new -key ' . $name[0] . '.key -out ' . $name[0] . '.csr -subj "/C=JP/ST=Tokyo/L=Chiyoda-ku/O=NTT COMMUNICATIONS CORPORATION/OU=NetworkServices/CN=sugakunn"');
}
// exec("openssl genrsa -des3 -out " . $name . " -passout pass:suga -rand sha1.dat 2048");


// $name = "kokoiti" . ".key";
// echo $name . "\n";

// $openssl = "openssl genrsa -des3 -out " . "server.key" . " -rand sha1.dat 2048";

// exec("openssl genrsa -des3 -out " . $name . " -rand sha1.dat 2048");

# openssl genrsa -des3 -out suga.key -passout pass:suga -rand sha1.dat 2048
# openssl rsa -passin pass:suga -in suga.key -out suga.key
# openssl req -new -key suga.key -out suga.csr -subj "/C=JP/ST=Tokyo/L=Chiyoda-ku/O=NTT COMMUNICATIONS CORPORATION/OU=NetworkServices/CN=sugakunn"

<?php
$fp = fopen('list.txt', 'r');

$str = '';

$option = array(
    0 => '/C=',
    1 => '/ST=',
    2 => '/L=',
    3 => '/O=',
    4 => '/OU=',
);

$getoption = array();

for ($i = 0; $i < 5; $i++) {
    $getoption[$i] = fgets($fp);
    $getoption[$i] = rtrim($getoption[$i], "\n");
    $str = $str . $option[$i] . $getoption[$i];
}
// echo $str . PHP_EOL;

while ($name = fgets($fp)) {
    $name = rtrim($name, "\n");
    $str2 = 'openssl req -new -key ' . $name . '.key -out ' . $name . '.csr -subj "' . $str . '/CN=' . $name . '"';
    exec('openssl genrsa -des3 -out ' . $name . '.key -passout pass:hoge -rand sha1.dat 2048');
    exec('openssl rsa -passin pass:hoge -in ' . $name . '.key -out ' . $name . '.key');
    exec($str2);
}

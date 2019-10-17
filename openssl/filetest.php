<?php
// $list = file(test);
$list = new SplFileObject('./test.txt');

// echo $list->fgets();

// echo $list->fgets();
$array = array();
$array[0] = 0;
$array[1] = 1;
echo $array[1];
foreach ($list as $key) {

}

$csr_option = "/C=JP/ST=Tokyo/L=Chiyoda-ku/O=NTT COMMUNICATIONS CORPORATION/OU=NetworkServices/CN=sugakunn"

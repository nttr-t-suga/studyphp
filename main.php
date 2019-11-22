<?php

// include ファイルがなくてもエラーにならない
// require ファイルがなかったらエラーになる
// require_once ファイルを一回だけ読み込む

// require_once('model.php');
// require_once('xmodel.php');

// $model = new Model('須賀');

// // $model->hello();

// $xmodel = new Xmodel('永野');

// $xmodel->hello();
// $xmodel->xhello();



// if ($model->test1 === '須賀') {
//     echo "{$model->test1}" . PHP_EOL . "須賀ですね";
// }

$test = array(
     'aaa' => array(
         'aika' => true,
         'aikb' => true,
     ),
     'bbb' => array(
         'bika' => true,
         'bikb' => true,
     ),
);
$ENV = null;
// $x = !empty;
if(!empty($x)){
    echo 1;
}

// echo $test['aaa']['aika'];
// if(isset($test['aaa']['aikb'])){
//     echo "ok";
// };
// echo array();
// echo isset($test);
// echo $test[$ENV][0];

// foreach ($test[$ENV] as $value) {
//     echo $value . PHP_EOL;
// }

// if ($ENV === 0) {
//     echo 1;
// }
// if (empty($test)) {
//     echo "empty";
// } elseif (!empty($test)) {
//     echo 'noempty';
// }

// if (strpos('abcd', 'ac')==0) {
//     echo 1;
// }
// echo strpos('abcd', 'cd');

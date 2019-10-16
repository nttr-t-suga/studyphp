<?php

// include ファイルがなくてもエラーにならない
// require ファイルがなかったらエラーになる
// require_once ファイルを一回だけ読み込む

require_once('model.php');
require_once('xmodel.php');

$model = new Model('須賀');

// $model->hello();

$xmodel = new Xmodel('永野');

$xmodel->hello();
// $xmodel->xhello();



// if ($model->test1 === '須賀') {
//     echo "{$model->test1}" . PHP_EOL . "須賀ですね";
// }

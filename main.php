<?php

// include ファイルがなくてもエラーにならない
// require ファイルがなかったらエラーになる
// require_once ファイルを一回だけ読み込む

require_once('model.php');

$model = new Model('須賀');

// $model2 = new model;


$model->hello();

// if ($model->test1 === '須賀') {
//     echo "{$model->test1}" . PHP_EOL . "須賀ですね";
// }

<?php
class Xmodel extends Model
{
    public $test2;

    public function __construct(string $word)
    {
        parent::__construct('ぱぱぱ');
        $this->test2 = $word;
    }

    public function hello()
    {
        echo $this->test2,'xmodelのハロー',PHP_EOL;
        parent::hello();
    }

    public function xhello()
    {
        echo $this->test2,'さんXこんにちは',PHP_EOL;
    }
}

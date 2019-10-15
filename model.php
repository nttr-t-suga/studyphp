<?php

class Model
{
    public $test1;

    public function __construct(string $test1)
    {
        $this->test1 = $test1;
    }

    public function hello()
    {
        echo $this->test1, 'さんこんにちは';
    }
}

<?php

class Test
{
    public $test1;


    public function __construct(string $test1)
    {
        $this->test1 = $test1;
    }

    public function hello()
    {
        echo $this->test1, 'さんこんにちは',PHP_EOL;
    }
}

class XTest extends Test
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



//オブジェクト指向テスト
interface Creature
{
    public function voice():string;
    public function laugh():string;
}


class Animal
{
    private $creature;

    public function __construct(Creature $creature)
    {
        $this->creature = $creature;
    }

    public function say()
    {
        echo $this->creature->voice(),PHP_EOL;
    }

    public function stroking(Type $var = null)
    {
        echo $this->creature->laugh(),PHP_EOL;
        echo $this->creature->test(100);
    }

}

class Cat implements Creature
{
    public function voice():string
    {
        return 'にゃぁ';
    }

    public function laugh():string
    {
        return 'ひげをいじる';
    }

    public function test(int $a):string
    {
        return 'aaa';
    }
}

class Dog implements Creature
{
    public function voice():string
    {
        return 'ワン';
    }

    public function laugh():string
    {
        return 'しっぽをふる';
    }
    
    public function test(string $a):int
    {
        return $a;
    }
}


$cat = new Animal(new Cat());
$cat->say();
$cat->stroking();

$pochi = new Animal(new Dog());
$pochi->say();
$pochi->stroking();

// $xtest = new XTest('須賀');
// $xtest->hello();

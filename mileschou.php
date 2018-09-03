<?php
// 'Hello! happy c9s!'

echo (new Arr())
    ->unshift('a')
    ->push('p')
    ->unshift('H')
    ->push('p')
    ->unshift('! ')
    ->unshift('lo')
    ->unshift('l')
    ->push('y')
    ->push(' ')
    ->push('c9s')
    ->unshift('He')
    ->push('!' . PHP_EOL)
;

class Arr
{
    private $d = [];

    public function push($data)
    {
        $this->d[] = $data;
        return $this;
    }

    public function unshift($data)
    {
        array_unshift($this->d, $data);
        return $this;
    }

    public function __toString()
    {
        return implode('', $this->d);
    }
}

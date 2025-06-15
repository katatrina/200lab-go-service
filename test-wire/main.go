package main

import "fmt"

type Foo struct {
	X int
}

func NewFoo() Foo {
	return Foo{X: 42}
}

type Bar struct {
	Foo Foo
}

func NewBar(foo Foo) Bar {
	return Bar{Foo: foo}
}

func (b Bar) Print() {
	fmt.Println(b.Foo.X)
}

func main() {
	bar := InitializeBar()
	bar.Print()
}

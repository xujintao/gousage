package main

import (
	"fmt"
)

type myI interface {
	foo()
}

type myS struct {
	age int
}

func (myS) foo() {
	fmt.Println("hi")
}

func main() {
	s := myS{20}
	var i myI = s
	i.foo() //https://github.com/derekparker/delve/issues/835
	//指针方法没问题，值方法需要单步两次
}

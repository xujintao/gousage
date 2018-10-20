package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
}

func main() {
	s1 := new(Student)
	s1.id = 100
	s1.name = "cat"
	s2 := Student{id: 200, name: "tom"}
	s3 := &Student{id: 300, name: "alice"}
	fmt.Println(s1, s2, s3) //&{100 cat} {200 tom} &{300 alice}
}

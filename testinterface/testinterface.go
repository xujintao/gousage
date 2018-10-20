package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name string
	age  int
}

func (u user) notify() {
	fmt.Println(u.name, u.age)
}

func main() {
	u := user{"bill", 24}
	i := notifier(&u)
	i.notify()
	fmt.Println(i) //&{bill 24}

	// pu := i.(*user)
	// pu.age = 20
	// fmt.Println(u) //{bill 20}
	// //new
	// vc := reflect.New(reflect.TypeOf(u))
	// fmt.Println(vc.Interface()) //&{ 0}
}

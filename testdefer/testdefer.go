package main

import "fmt"

//main
func main() {
	defer fmt.Println("dmain")

	a()
	fmt.Println("main")
}

//a
func a() {
	defer fmt.Println("da")
	b()
	fmt.Println("a")
}

//b
func b() {
	defer fmt.Println("db")
	c()
	fmt.Println("b")
}

//c
func c() {
	defer fmt.Println("dc")

	fmt.Println("c")
	// panic("test")
}

package main

import (
	"fmt"
)

type Animal struct {
	name string
	age  int
}

func (this *Animal) Eat() {
	fmt.Printf("Animal::Eat\n")
}
func (this *Animal) Sleep() {
	fmt.Printf("Animal::Sleep\n")
}

type Human struct {
	Animal
	ext int
}

func (this *Human) Eat() {
	fmt.Println("Human::Eat")
}

func main() {
	//Human
	alice := Human{Animal{"alice", 20}, 1}
	alice.Eat()
	alice.Sleep()
	fmt.Printf("%s, %d\n", alice.name, alice.age)
}

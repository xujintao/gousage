package main

import (
	"fmt"
)

func main() {
	s := []int{10, 20, 30, 40}
	for _, v := range s[1:] {
		fmt.Println(v)
	}
}

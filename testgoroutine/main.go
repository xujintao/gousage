package main

import (
	"fmt"
)

func main() {
	c1 := make(chan bool, 1)
	c2 := make(chan bool) //无缓冲channel的读写没有意义
	c3 := make(chan bool)
	done := make(chan bool)

	//g1
	go func() {
		for i := 0; i < 20; i++ {
			<-c1
			fmt.Print("1")
			c2 <- true
		}
	}()

	//g2
	go func() {
		for i := 0; i < 20; i++ {
			<-c2
			fmt.Print("2")
			c3 <- true
		}
	}()

	//g3
	go func() {
		for i := 0; i < 20; i++ {
			<-c3
			fmt.Print("3 ")
			c1 <- true
		}
		close(done)
	}()

	c1 <- true
	<-done
	fmt.Println("exit")
}

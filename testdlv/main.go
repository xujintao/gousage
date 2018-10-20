package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("[%s:%d] ", "g1", i)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("[%s:%d] ", "g2", i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("end")
}

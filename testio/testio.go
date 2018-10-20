package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("G:/game/mu.txt")
	if err != nil {
		fmt.Println(err)
	}

	// bfr := bufio.NewReader(strings.NewReader(""))

	bufr := bufio.NewReader(f)
	b5, err := bufr.Peek(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b5))
	f.Close()
}

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h2 := sha1.New()
	b := []byte("123")
	h2.Write(b)
	h2.Write([]byte("abc"))
	b2 := h2.Sum(nil)
	b3 := h2.Sum(b)
	fmt.Println(b2)
	fmt.Println(b3)
}

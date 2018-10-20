package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().UTC().Format(time.RFC3339)) //2018-10-07T16:38:28Z
	t, _ := time.Parse("2006-01-02T15:04:05Z0700", "2018-10-07T16:38:28Z")
	fmt.Println(t)
}

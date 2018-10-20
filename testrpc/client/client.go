package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Params struct {
	Width  int
	Height int
}

func main() {
	// rpc, err := rpc.DialHTTP("tcp", "192.168.6.201:8080")//http-rpc
	rpc, err := rpc.Dial("tcp", "192.168.6.201:8080") //tcp-rpc
	// rpc, err := jsonrpc.Dial("tcp", "192.168.6.201:8080") //json-rpc
	if err != nil {
		log.Fatal(err)
	}

	ret := 0
	if err := rpc.Call("Rect.Area", Params{50, 100}, &ret); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)

	ret = 0
	if err := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)

	time.Sleep(2 * time.Second)
	rpc.Close()
	time.Sleep(20 * time.Second)
}

package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//参数
type Params struct {
	Width  int
	Height int
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//---1.http-rpc
// func main() {
// 	rect := new(Rect)
// 	rpc.Register(rect)
// 	rpc.HandleHTTP()
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// ---2.tcp-rpc
// func main() {
// 	rect := new(Rect)
// 	rpc.Register(rect)
// 	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8080")
// 	checkError(err)
// 	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
// 	checkError(err)
// 	for {
// 		conn, err := tcpListen.Accept()
// 		if err != nil {
// 			continue
// 		}
// 		go rpc.ServeConn(conn)
// 	}
// }

//---3.json-rpc
func main() {
	rpc.Register(&Rect{})
	//监听1
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8080")
	// checkError(err)
	// tcpListen, err := net.ListenTCP("tcp", tcpAddr)

	//监听2
	tcpListen, err := net.Listen("tcp", ":8080")
	checkError(err)
	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

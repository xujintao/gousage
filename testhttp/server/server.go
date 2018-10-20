package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//读body

	//方式一，直接读body串
	// // body, err := req.GetBody()//为什么不行？
	// data := make([]byte, req.ContentLength)
	// nLen, err := req.Body.Read(data)
	// if err != io.EOF {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(data), nLen)
	// rw.Write([]byte(data))

	//方式二，form
	req.ParseForm()
	// name := req.Form.Get("name")//GET方法querystring
	name := req.PostForm.Get("name") //POST方法body
	fmt.Println(name)
	rw.Write([]byte(name))
}

func main() {
	endRunning := make(chan bool)
	// http.ListenAndServe(":8080", &myHandler{}) //为什么不这样用？
	server := &http.Server{Addr: ":8080", Handler: &myHandler{}}
	go func() {
		server.ListenAndServe()
		endRunning <- true
	}()
	<-endRunning
	// time.Sleep(10 * time.Second)
	server.Close()
	close(endRunning)
	fmt.Println("close server")
}

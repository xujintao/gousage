package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// application/x-www-form-urlencoded
	// text/xml
	// application/json
	// multipart/form-data
	// r, err := http.Post("http://127.0.0.1:8080/testgin/urlencode?id=123&page=1", "application/x-www-form-urlencoded",
	// 	strings.NewReader("username=alice&password=1234"))
	// r, err := http.Post("http://127.0.0.1:8080/testgin/json", "application/json",
	// 	strings.NewReader(`{"id":"123","page":"1","username":"alice","password":1234}`))

	tp := http.DefaultTransport.(*http.Transport)
	tp.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //不验证证书
	tp.TLSHandshakeTimeout = 100000 * time.Second              //调试用
	client := &http.Client{
		//连接池配置
		Transport: tp,

		//会话超时时间
		Timeout: 10 * time.Second,
	}

	//第一次会话
	// req, err := http.NewRequest("GET", "http://192.168.6.201:8080/r1", nil)
	req, err := http.NewRequest("GET", "https://192.168.6.201:8080/hellohttps", nil)
	if err != nil {
		log.Fatal(err)
	}
	r, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	var b bytes.Buffer
	b.Write([]byte("hi"))

	// data := make([]byte, 10)
	// n, err := r.Body.Read(data)
	// fmt.Printf("data:%s, num:%d", data, n)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(data)
	fmt.Printf("%s\n", data)
	fmt.Println(string(data))

	log.Print(data)
	log.Printf("%s\n", data)
	log.Println(string(data))

	//第二次会话
	// req, err = http.NewRequest("GET", "http://192.168.6.201:8080/r2", nil)
	req, err = http.NewRequest("GET", "https://192.168.6.201:8080/hellohttps", nil)
	if err != nil {
		log.Fatal(err)
	}
	r, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

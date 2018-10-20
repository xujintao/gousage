package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type User struct {
	Id       string `form:"id" json:"id"`
	Page     string `form:"page" json:"page"`
	Username string `form:"username" json:"username"`
	Password int    `form:"password" json:"password"`
}

func main() {
	http.HandleFunc("/r1", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello, r1"))
	})
	http.HandleFunc("/r2", func(rw http.ResponseWriter, req *http.Request) {
		io.WriteString(rw, "hello, r2")
	})

	/*
		POST /testgin/urlencode?id=1234&page=1 HTTP/1.1
		Content-Type: application/x-www-form-urlencoded

		username=alice&password=1234
	*/
	http.HandleFunc("/testgin/urlencode", func(rw http.ResponseWriter, req *http.Request) {
		var u User
		req.ParseForm()
		u.Id = req.Form.Get("id")
		u.Page = req.Form.Get("page")
		u.Username = req.PostForm.Get("username")
		u.Password, _ = strconv.Atoi(req.PostForm.Get("password"))
		fmt.Println(u)
		data := make(url.Values) //url编码
		data.Set("id", u.Id)
		data.Set("page", u.Page)
		data.Set("username", u.Username)
		data.Set("password", strconv.Itoa(u.Password))
		rw.Write([]byte(data.Encode()))
	})

	/*
		POST /testgin/json HTTP/1.1
		Content-Type: application/json;charset=utf-8

		{"id":"123","page":"1","username":"alice","password":1234}
	*/
	http.HandleFunc("/testgin/json", func(rw http.ResponseWriter, req *http.Request) {
		var u User
		// json.Unmarshal(req.Body, &u)
		json.NewDecoder(req.Body).Decode(&u)
		fmt.Println(u)
		data, _ := json.Marshal(&u) //序列化
		rw.Write(data)
	})

	log.Println("Listening and serving HTTP on :8080")
	http.ListenAndServe(":8080", nil)
}

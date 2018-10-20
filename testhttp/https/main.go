package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hellohttps", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello, https"))
	})

	log.Println("Listening and serving HTTPS on :8080")
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}

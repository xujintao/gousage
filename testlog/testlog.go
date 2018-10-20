package main

import (
	"io"
	"log"
	"os"
)

var Error *log.Logger

func init() {
	//设置log.std
	log.SetPrefix("INFO: ")
	log.SetFlags(log.LstdFlags)

	//设置Error
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.LstdFlags)
}

func main() {
	bRet := 1 > 2
	log.Printf("%d %s %t\n", 1, "hello", bRet)
	Error.Println("hello world")
}

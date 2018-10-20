package main

import (
	"flag"

	"github.com/golang/glog"
)

func init() {
	flag.Set("log_dir", ".")
	flag.Set("v", "3")
	flag.Set("alsologtostderr", "true")
	flag.Parse()
}

func main() {

	glog.Info("hello, info")
	glog.Warning("hello, warning")
	glog.Error("hello, error")
	glog.Fatal("hello, fatal")
	glog.Flush()
}

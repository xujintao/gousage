package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//测试os.Args
	var appPath string
	var err error
	if appPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	fmt.Println(appPath, err)

	//测试Join方法
	workPath, err := os.Getwd()
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	fmt.Println(appConfigPath)

	//测试abs
	absConfigPath, err := filepath.Abs("conf/app2.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println(absConfigPath)
}

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/glog"
	yaml "gopkg.in/yaml.v2"
)

// type Config struct {
// 	DB struct {
// 		Pgsql struct {
// 			Name     string `yaml:"db.pgsql.name"`
// 			User     string `yaml:"db.pgsql.user"`
// 			Password string `yaml:"db.pgsql.password"`
// 			Host     string `yaml:"db.pgsql.host"`
// 			Port     uint   `yaml:"db.pgsql.port"`
// 		}
// 		Redis struct {
// 			Host     string `yaml:"db.redis.host"`
// 			Port     string `yaml:"db.redis.port"`
// 			UserName string `yaml:"db.redis.username"`
// 			PassWord string `yaml:"db.redis.password"`
// 		}
// 	}
// }

type Config struct {
	DB struct {
		Pgsql struct {
			Name     string
			User     string
			Password string
			Host     string
			Port     uint
		}
		Redis struct {
			Host     string
			Port     string
			UserName string
			PassWord string
		}
	}
}

var Addr = "."

func main() {
	yamlFile, err := ioutil.ReadFile("./demo.yaml")
	if err != nil {
		glog.Exitf("加載配置文件失敗 %s ", err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		glog.Exitf("格式化配置文件失敗,%s ", err.Error())
	}
	fmt.Print(cfg)
}

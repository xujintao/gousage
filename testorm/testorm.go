package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql" //github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
)

//点赞表
type Like struct {
	gorm.Model
	//ID    int    `gorm:"primary_key"`
	//Date  time.Time
	User  string `gorm:"type:varchar(20); not null;index:idx_user_title"`
	Title string //`gorm:"type:varchar(128);not null;index:idx_user_title"`
}

func (t *Like) TableName() string {
	return "t_like"
}

func main() {
	//数据库连接
	db, err := gorm.Open("mysql", "root:1234@tcp(192.168.6.200:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)

	//表结构操作（不用orm）
	//创建表
	if !db.HasTable(&Like{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			panic(err)
		}
	}
	//删除表，修改表，增删改索引，外键等

	//CURD
	//1, insert
	like := &Like{
		User:  "alice",
		Title: "有点尴尬了 我太单纯了",
	}
	if err := db.Create(like).Error; err != nil {
		fmt.Println(err)
	}

	for {
		time.Sleep(3 * time.Second)
	}
}

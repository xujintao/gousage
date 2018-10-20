package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type JsonController struct {
	beego.Controller
}

func (c *JsonController) Post() {
	type user struct {
		Username string `json:"username"`
		Passwd   int    `json:"password"`
	}

	type resultData struct {
		Uid    string `json:"uid"`
		Result string `json:"result"`
	}

	//原生方式
	req := c.Ctx.Request
	rw := c.Ctx.ResponseWriter.ResponseWriter
	var u user
	// json.Unmarshal(req.Body, &u)
	json.NewDecoder(req.Body).Decode(&u)
	fmt.Println(u.Username, u.Passwd)
	r := &resultData{"1001", "success"} //映射数据库结构体
	data, _ := json.Marshal(r)          //序列化
	rw.Write(data)

	//beego方式
	// var u user
	// json.Unmarshal(ctx.Input.CopyBody(50), &u)
	// r := &resultData{"1001", "success"} //映射数据库结构体
	// ctx.Output.JSON(r, false, false)
}

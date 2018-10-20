package controllers

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/astaxie/beego"
)

type URLEncodeController struct {
	beego.Controller
}

func (c *URLEncodeController) Post() {
	//原生方式，querystring
	req := c.Ctx.Request
	rw := c.Ctx.ResponseWriter.ResponseWriter
	req.ParseForm()
	username := req.PostForm.Get("username")
	password, _ := strconv.Atoi(req.PostForm.Get("passwd"))
	fmt.Println(username, password)
	data := make(url.Values) //url编码
	data.Set("uid", "1001")
	data.Add("result", "success")
	rw.Write([]byte(data.Encode()))

}

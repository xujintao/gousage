package main

import (
	"test/testbeego/controllers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/testgin/urlencode", &controllers.URLEncodeController{})
	beego.Router("/testgin/json", &controllers.JsonController{})
	beego.Router("/testbeego/urlencode", &controllers.URLEncodeController{})
	beego.Router("/testbeego/json", &controllers.JsonController{})
	beego.Router("/hello", &controllers.JsonController{})
	beego.Run()
}

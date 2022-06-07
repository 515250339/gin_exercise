package routers

import (
	"beegogorm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/add/", &controllers.UserController{}, "get:UserAdd")
	beego.Router("/user/delete/", &controllers.UserController{}, "get:UserDelete")
	beego.Router("/user/update/", &controllers.UserController{}, "get:UserUpdate")
	beego.Router("/user/get/", &controllers.UserController{}, "get:GetUser")
}

package main

import (
	"beegogorm/models"
	_ "beegogorm/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	defer models.DB.Close() //关闭数据库连接
}
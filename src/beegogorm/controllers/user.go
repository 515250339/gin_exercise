package controllers

import (
	"beegogorm/models"
	"github.com/astaxie/beego"
	"time"
)

type UserController struct {
	beego.Controller
}

// UserAdd 增加用户
func (c *UserController) UserAdd() {
	user := models.User{
		UserName: "zs",
		Age:      22,
		Email:    "email@email",
		AddTime:  int(time.Now().Unix()),
	}
	models.DB.Create(&user)
	c.Ctx.WriteString("添加数据成功")
}

// UserDelete 删除用户
func (c *UserController) UserDelete() {
	user := models.User{Id: 1}
	models.DB.Delete(&user)
	c.Ctx.WriteString("删除成功")
}

// UserUpdate 更新用户
func (c *UserController) UserUpdate() {
	// 查找id=9的数据
	user := models.User{Id: 9}
	// models.DB.First(&user)

	// 执行修改
	user.UserName = "秀儿"
	models.DB.Save(&user)

	c.Ctx.WriteString("修改成功")
}

// GetUser 获取用户
func (c *UserController) GetUser() {
	// 查询单个
	//user := models.User{UserName: "秀儿"}
	//models.DB.Find(&user)

	// 查询所有
	var user []models.User
	models.DB.Find(&user)
	c.Data["json"] = user
	c.ServeJSON()
}

package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	//和数据库建立连接
	DB, err = gorm.Open("mysql", "root:root@/beego_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		beego.Error()
	}
	if !DB.HasTable(&User{}) {
		// 创建表
		DB.CreateTable(&User{})                                            // 根据User结构体建表
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}) // 设置表结构的存储引擎为InnoDB
	}
}

package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Core() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
	}
	// 创建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil
	} // 根据User结构体建表
	fmt.Println("建立连接成功")
	return db
}

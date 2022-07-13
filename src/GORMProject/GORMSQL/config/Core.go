package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Core() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println(err)
	}
	// 创建表
	err = db.AutoMigrate(&User{})
	err = db.AutoMigrate(&CreditCard{})
	err = db.AutoMigrate(&UserInfo{})
	err = db.AutoMigrate(&DefaultTest1{})
	err = db.AutoMigrate(&DefaultTest2{})
	err = db.AutoMigrate(&TestModel{})
	//err = db.AutoMigrate(&DefaultTest3{})
	if err != nil {
		return nil
	} // 根据User结构体建表
	fmt.Println("建立连接成功")
	return db
}

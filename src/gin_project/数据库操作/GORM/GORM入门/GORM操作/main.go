package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	Id       int
	UserName string
	Age      int
	Email    string
	AddTime  int
}

// TableName 定义结构体操作的数据库表
func (User) TableName() string {
	return "user"
}

type Postparams1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"`
}

type Postparams2 struct {
	Name string `json:"name" uri:"name"`
	Age  int    `json:"age" uri:"age"`
	Sex  bool   `json:"sex" uri:"sex"`
}

type Postparams3 struct {
	Name string `json:"name" form:"name" binding:"required"`
	Age  int    `json:"age" form:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex" form:"sex" binding:"required"`
}

var DB *gorm.DB
var err error

func init() {
	// 和数据库建立连接
	DB, err = gorm.Open("mysql", "root:root@/gin_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	if !DB.HasTable(&User{}) {
		// 创建表
		DB.CreateTable(&User{})
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}) // 设置表结构的存储引擎为InnoDB

	}
}

func UserAdd(UserName string, Age int, Email string) string {
	user := User{
		UserName: UserName,
		Age:      Age,
		Email:    Email,
		AddTime:  int(time.Now().Unix()),
	}
	DB.Create(&user)
	return "ok"
}

func main() {
	UserName := "zs"
	Age := 10
	Email := "zs.email"
	UserAdd(UserName, Age, Email)
	defer DB.Close()
}

//func main() {
//
//	r := gin.Default() // 携带基础中间件启动
//	// GET 请求
//	r.GET("/path/:id", func(c *gin.Context) {
//		// 获取 url 参数
//		UserName := c.Param("UserName")
//		// 获取 params 参数，默认为 wang
//		//Age := c.DefaultQuery("Age", "test")
//		// 获取 params 参数
//		Email := c.Query("Email")
//		UserAdd(UserName, 1, Email)
//		c.JSON(200, gin.H{})
//	})
//	// POST 请求
//	r.POST("/path", func(c *gin.Context) {
//		// 获取请求体中参数 默认为 wang
//		user := c.DefaultPostForm("user", "wang")
//		// 获取请求体中参数
//		pwd := c.PostForm("pwd")
//		c.JSON(200, gin.H{
//			"user": user,
//			"pwd":  pwd,
//		})
//	})
//	// PUT 请求
//	r.PUT("/path", func(c *gin.Context) {
//		user := c.DefaultPostForm("user", "wang")
//		pwd := c.PostForm("pwd")
//		c.JSON(200, gin.H{
//			"user": user,
//			"pwd":  pwd,
//		})
//	})
//	// DELETE 请求
//	r.DELETE("/path/:id", func(c *gin.Context) {
//		id := c.Param("id")
//		user := c.DefaultPostForm("user", "wang")
//		pwd := c.PostForm("pwd")
//		c.JSON(200, gin.H{
//			"id":   id,
//			"user": user,
//			"pwd":  pwd,
//		})
//	})
//	err := r.Run()
//	if err != nil {
//		return
//	} // listen and serve on 0.0.0.0:8080
//	defer DB.Close() //关闭数据库连接
//}

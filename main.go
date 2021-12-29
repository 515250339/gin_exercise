package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

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

func mustBig(f1 validator.FieldLevel) bool {
	if f1.Field().Interface().(int) <= 18 {
		return false
	}
	return true
}

func main() {
	r := gin.Default() // 携带基础中间件启动
	// 参数校验 age 是否大于 18
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustBig", mustBig)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	// 通过 ShouldBindQuery 验证参数
	r.POST("/testBind3", func(c *gin.Context) {
		var p Postparams3
		err := c.ShouldBindQuery(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":   "失败了",
				"data":  gin.H{},
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功了",
				"data": p,
			})
		}
	})
	// 通过 ShouldBindUri 验证参数
	r.POST("/testBind2/:name/:age/:sex", func(c *gin.Context) {
		var p Postparams2
		err := c.ShouldBindUri(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "失败了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功了",
				"data": p,
			})
		}
	})
	// 通过 ShouldBindJSON 验证参数
	r.POST("/testBind1", func(c *gin.Context) {
		var p Postparams1
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "失败了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功了",
				"data": p,
			})
		}
	})
	// GET 请求
	r.GET("/path/:id", func(c *gin.Context) {
		// 获取 url 参数
		id := c.Param("id")
		// 获取 params 参数，默认为 wang
		user := c.DefaultQuery("user", "wang")
		// 获取 params 参数
		pwd := c.Query("pwd")

		c.JSON(200, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})
	// POST 请求
	r.POST("/path", func(c *gin.Context) {
		// 获取请求体中参数 默认为 wang
		user := c.DefaultPostForm("user", "wang")
		// 获取请求体中参数
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	// PUT 请求
	r.PUT("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "wang")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	// DELETE 请求
	r.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultPostForm("user", "wang")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}

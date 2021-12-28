package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // 携带基础中间件启动
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
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

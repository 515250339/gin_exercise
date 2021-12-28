package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // 携带基础中间件启动
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "wang")
		pwd := c.Query("pwd")

		c.JSON(200, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "wang")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	r.PUT("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "wang")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
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

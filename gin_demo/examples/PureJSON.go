package main

import "github.com/gin-gonic/gin"

func PureJSON() {
	r := gin.Default()

	// 提供unicode实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello World</b>",
		})
	})

	// 提供字面字符
	r.GET("pursejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello World</b>",
		})
	})

	r.Run()
}

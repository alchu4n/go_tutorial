package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AsciiJSON() {
	r := gin.Default()
	r.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

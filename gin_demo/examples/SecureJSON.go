package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
func SecureJSON() {
	r := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")
	r.GET("/someJson", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONP() {
	r := gin.Default()
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{"foo": "bar"}
		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	r.Run()
}

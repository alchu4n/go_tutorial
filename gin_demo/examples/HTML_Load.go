package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadHTMLGlob() {
	r := gin.Default()
	//使用 LoadHTMLGlob() 或者 LoadHTMLFiles()
	r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.Run(":8080")
}

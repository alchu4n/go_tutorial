package main

import "github.com/gin-gonic/gin"

func Without_Middleware() {
	//不使用默认的中间件
	r := gin.New()
	// Default 使用 Logger 和 Recovery 中间件
	//r := gin.Default()

	r.Run()
}

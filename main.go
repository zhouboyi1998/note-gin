package main

import "github.com/gin-gonic/gin"

func main() {
	// 使用 Gin 默认路由引擎
	route := gin.Default()
	// GET 请求, 路径 /ping, 返回值内容为 {"message": "ok"}
	route.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "ok"})
	})
	// 监听 127.0.0.1:18091
	route.Run("127.0.0.1:18091")
}

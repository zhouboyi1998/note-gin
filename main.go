package main

import "github.com/gin-gonic/gin"

func main() {
	// 新建 Gin 实例
	app := gin.Default()
	// GET 请求, 路径 /ping, 返回值内容为 {"message": "ok"}
	app.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "ok"})
	})
	// 启动服务
	app.Run("127.0.0.1:18091")
}

package main

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/application"
	"note-gin/src/router"
)

func main() {
	// 新建 Gin 实例
	app := gin.Default()
	// 注册路由
	router.RegisterRouter(app)
	// 启动服务
	app.Run(application.App.Server.Host + ":" + application.App.Server.Port)
}

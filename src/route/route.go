package route

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/controller"
)

// Route 自定义路由
func Route() *gin.Engine {
	// 使用 Gin 框架默认路由
	r := gin.Default()

	// 添加命令相关路由组
	commandGroup := r.Group("/note/command")
	// 添加路由规则
	commandGroup.GET("/one/:commandName", controller.One)

	return r
}

package route

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/controller"
)

// Route 自定义路由规则
func Route() *gin.Engine {
	// 使用 Gin 默认路由
	r := gin.Default()

	// 添加 Linux 命令相关路由组
	commandGroup := r.Group("/note/command")
	// 添加 Linux 命令相关路由规则
	commandGroup.GET("/one/:commandName", controller.One)
	commandGroup.GET("/list", controller.List)
	commandGroup.GET("/list/name", controller.ListName)
	commandGroup.POST("/insert", controller.InsertOne)
	commandGroup.POST("/insert/many", controller.InsertMany)
	commandGroup.PUT("/update", controller.UpdateOne)
	commandGroup.DELETE("/delete/:commandId", controller.DeleteOne)

	return r
}

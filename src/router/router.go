package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/controller"
)

// Route 自定义路由规则
func Route() *gin.Engine {
	// 使用 Gin 默认路由
	r := gin.Default()

	// 添加 /command 路由组
	g := r.Group("/note/command")
	// 添加 /command 相关路由规则
	g.GET("/one/:commandName", controller.One)
	g.GET("/list", controller.List)
	g.GET("/list/name", controller.ListName)
	g.POST("/insert", controller.InsertOne)
	g.POST("/insert/many", controller.InsertMany)
	g.PUT("/update", controller.UpdateOne)
	g.DELETE("/delete/:commandId", controller.DeleteOne)

	return r
}

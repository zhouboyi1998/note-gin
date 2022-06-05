package router

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/controller"
)

// RegisterRouter 注册路由
func RegisterRouter(e *gin.Engine) {
	// 新建命令路由组
	g := e.Group("/command")
	{
		// 添加命令相关路由
		g.GET("/:commandId", controller.One)
		g.GET("/", controller.List)
		g.POST("/", controller.Insert)
		g.POST("/batch", controller.InsertBatch)
		g.PUT("/", controller.Update)
		g.PUT("/batch", controller.UpdateBatch)
		g.DELETE("/:commandId", controller.Delete)
		g.DELETE("/batch", controller.DeleteBatch)
		g.GET("/select/:commandName", controller.Select)
		g.GET("/name-list", controller.NameList)
	}
}

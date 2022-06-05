package controller

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/mongo"
)

// One 按命令名称查询一条命令
func One(c *gin.Context) {
	command := mongo.One(c.Param("commandName"), c)
	c.JSON(200, command)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/mongo"
)

// One 按 Linux 命令名称查询一条命令
func One(c *gin.Context) {
	command := mongo.One(c, c.Param("commandName"))
	c.JSON(200, command)
}

func List(c *gin.Context) {
	commandList := mongo.List(c)
	c.JSON(200, commandList)
}

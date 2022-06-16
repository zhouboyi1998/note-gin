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

// List 查询所有 Linux 命令
func List(c *gin.Context) {
	commandList := mongo.List(c)
	c.JSON(200, commandList)
}

// ListName 查询所有 Linux 命令的名称
func ListName(c *gin.Context) {
	nameList := mongo.ListName(c)
	c.JSON(200, nameList)
}

// InsertOne 插入单条 Linux 命令
func InsertOne(c *gin.Context) {
	result, commandName := mongo.InsertOne(c)
	c.JSON(200, gin.H{
		"Id":      result,
		"Command": commandName,
	})
}

// InsertMany 插入多条 Linux 命令
func InsertMany(c *gin.Context) {
	result := mongo.InsertMany(c)
	c.JSON(200, result)
}

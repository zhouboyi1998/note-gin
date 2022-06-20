package controller

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/mongo"
)

// One 按 Linux 命令名称查询一条命令
func One(c *gin.Context) {
	commandName := mongo.One(c, c.Param("commandName"))
	c.JSON(200, commandName)
}

// List 查询所有 Linux 命令
func List(c *gin.Context) {
	commandArray := mongo.List(c)
	c.JSON(200, commandArray)
}

// ListName 查询所有 Linux 命令的名称
func ListName(c *gin.Context) {
	nameArray := mongo.ListName(c)
	c.JSON(200, nameArray)
}

// InsertOne 插入单条 Linux 命令
func InsertOne(c *gin.Context) {
	result, commandName := mongo.InsertOne(c)
	c.JSON(200, gin.H{
		"result":  result,
		"command": commandName,
	})
}

// InsertMany 插入多条 Linux 命令
func InsertMany(c *gin.Context) {
	result := mongo.InsertMany(c)
	c.JSON(200, result)
}

// UpdateOne 更新单条 Linux 命令
func UpdateOne(c *gin.Context) {
	result := mongo.UpdateOne(c)
	c.JSON(200, result)
}

// DeleteOne 删除单条 Linux 命令
func DeleteOne(c *gin.Context) {
	result, objectId := mongo.DeleteOne(c, c.Param("commandId"))
	c.JSON(200, gin.H{
		"result": result,
		"_id":    objectId,
	})
}

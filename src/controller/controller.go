package controller

import (
	"github.com/gin-gonic/gin"
	"note-gin/src/mongo"
)

// One 按 ObjectId 查询单条 Linux 命令
func One(c *gin.Context) {
	commandId := c.Param("commandId")
	command := mongo.One(c, commandId)
	c.JSON(200, command)
}

// OneByName 按 Linux 命令名称查询单条 Linux 命令
func OneByName(c *gin.Context) {
	commandName := c.Param("commandName")
	command := mongo.OneByName(c, commandName)
	c.JSON(200, command)
}

// List 查询所有 Linux 命令
func List(c *gin.Context) {
	commandArray := mongo.List(c)
	c.JSON(200, commandArray)
}

// ListByName 查询所有 Linux 命令的名称
func ListByName(c *gin.Context) {
	nameArray := mongo.ListByName(c)
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
	commandId := c.Param("commandId")
	result, objectId := mongo.DeleteOne(c, commandId)
	c.JSON(200, gin.H{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteMany 删除多条 Linux 命令
func DeleteMany(c *gin.Context) {
	result, objectIds := mongo.DeleteMany(c)
	c.JSON(200, gin.H{
		"result": result,
		"_ids":   objectIds,
	})
}

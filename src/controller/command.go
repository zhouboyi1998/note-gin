package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"note-gin/src/repository"
)

// One 根据id查询命令
func One(c *gin.Context) {
	command := repository.One(c)
	c.JSON(http.StatusOK, command)
}

// List 查询命令列表
func List(c *gin.Context) {
	commandArray := repository.List(c)
	c.JSON(http.StatusOK, commandArray)
}

// Insert 新增命令
func Insert(c *gin.Context) {
	result, commandName := repository.Insert(c)
	c.JSON(http.StatusOK, gin.H{
		"result":  result,
		"command": commandName,
	})
}

// InsertBatch 批量新增命令
func InsertBatch(c *gin.Context) {
	result := repository.InsertBatch(c)
	c.JSON(http.StatusOK, result)
}

// Update 修改命令
func Update(c *gin.Context) {
	result := repository.Update(c)
	c.JSON(http.StatusOK, result)
}

// UpdateBatch 批量修改命令
func UpdateBatch(c *gin.Context) {
	result := repository.UpdateBatch(c)
	c.JSON(http.StatusOK, result)
}

// Delete 删除命令
func Delete(c *gin.Context) {
	result, objectId := repository.Delete(c)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteBatch 批量删除命令
func DeleteBatch(c *gin.Context) {
	result, objectIds := repository.DeleteBatch(c)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"_ids":   objectIds,
	})
}

// Select 查询命令
func Select(c *gin.Context) {
	command := repository.Select(c)
	c.JSON(http.StatusOK, command)
}

// NameList 查询命令名称列表
func NameList(c *gin.Context) {
	nameArray := repository.NameList(c)
	c.JSON(http.StatusOK, nameArray)
}

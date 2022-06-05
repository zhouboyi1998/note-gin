package mongo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"note-gin/src/model"
)

// Connection 连接 MongoDB
func Connection(c *gin.Context) *mongo.Collection {
	// 连接 MongoDB 客户端
	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://root:123456@127.0.0.1:27017/"))
	if err != nil {
		err.Error()
	}

	// 连接 note 数据库, linux_command 表
	collection := client.Database("note").Collection("linux_command")

	return collection
}

// One 按命令名称查询一条命令
func One(commandName string, c *gin.Context) model.Command {
	// 获取数据库连接
	collection := Connection(c)

	// 按命令名称查询数据
	result := collection.FindOne(c, bson.M{
		"command": commandName,
	})

	// 转换为 Command 结构体
	command := model.Command{}
	err := result.Decode(&command)
	if err != nil {
		err.Error()
	}

	return command
}

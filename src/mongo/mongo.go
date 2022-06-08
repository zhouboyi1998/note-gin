package mongo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"note-gin/src/model"
)

// Connection 连接 MongoDB
func Connection(c *gin.Context) *mongo.Collection {
	// 连接 MongoDB 客户端
	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://root:123456@host.docker.internal:27017/"))
	if err != nil {
		log.Println(err)
	}

	// 连接 note 数据库, linux_command 表
	collection := client.Database("note").Collection("linux_command")

	return collection
}

// One 按命令名称查询一条命令
func One(c *gin.Context, commandName string) model.Command {
	// 获取数据库连接
	collection := Connection(c)

	// 按 Linux 命令名称查询数据
	result := collection.FindOne(c, bson.M{
		"command": commandName,
	})

	// 转换为结构体
	command := model.Command{}
	err := result.Decode(&command)
	if err != nil {
		log.Println(err)
	}

	return command
}

// List 查询所有 Linux 命令
func List(c *gin.Context) []model.Command {
	// 获取数据库连接
	collection := Connection(c)

	// 查询所有 Linux 命令
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 转换为结构体数组
	var commandList []model.Command
	// 返回值 cursor 相当于一个指针, 需要 Next() 遍历一个一个获取数据
	for cursor.Next(c) {
		command := model.Command{}
		cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		commandList = append(commandList, command)
	}

	return commandList
}

// ListName 查询所有 Linux 命令的名称
func ListName(c *gin.Context) []string {
	// 获取数据库连接
	collection := Connection(c)

	// 查询所有 Linux 命令
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 转换为结构体数组
	var nameList []string
	// 返回值 cursor 相当于一个指针, 需要 Next() 遍历一个一个获取数据
	for cursor.Next(c) {
		command := model.Command{}
		cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		nameList = append(nameList, command.Command)
	}

	return nameList
}

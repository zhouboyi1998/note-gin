package mongo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"note-gin/src/conf"
	"note-gin/src/model"
)

// Connection 连接 MongoDB
func Connection(c *gin.Context) *mongo.Collection {
	// 从配置文件中读取连接配置
	uri := "mongodb://" + conf.Config.Mongo.Username + ":" + conf.Config.Mongo.Password + "@" + conf.Config.Mongo.Host + ":" + conf.Config.Mongo.Port + "/"
	// 连接 MongoDB 客户端
	client, err := mongo.Connect(c, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}

	// 连接配置文件指定的数据库和文档集合
	collection := client.Database(conf.Config.Mongo.Database).Collection(conf.Config.Mongo.Collection)

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

	// 结构体对象
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

	// 结构体数组
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

	// 字符串数组
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

// InsertOne 插入单条 Linux 命令
func InsertOne(c *gin.Context) (*mongo.InsertOneResult, string) {
	// 获取数据库连接
	collection := Connection(c)

	// 结构体对象
	command := model.Command{}
	// 将请求体参数赋值到结构体对象上
	errBind := c.ShouldBind(&command)
	if errBind != nil {
		log.Println(errBind)
	}
	// 生成 ObjectID
	command.Id = primitive.NewObjectID()

	// 插入单条 Linux 命令
	result, err := collection.InsertOne(c, command)
	if err != nil {
		log.Println(err)
	}

	return result, command.Command
}

// InsertMany 插入多条 Linux 命令
func InsertMany(c *gin.Context) *mongo.InsertManyResult {
	// 获取数据库连接
	collection := Connection(c)

	// interface 数组
	var commandList []interface{}

	// 将请求体参数赋值到结构体对象上
	errBind := c.ShouldBind(&commandList)
	if errBind != nil {
		log.Println(errBind)
	}

	// 插入多条 Linux 命令
	result, err := collection.InsertMany(c, commandList)
	if err != nil {
		log.Println(err)
	}

	return result
}

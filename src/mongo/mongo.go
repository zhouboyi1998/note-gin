package mongo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"note-gin/src/application"
	"note-gin/src/model"
)

// Collection 连接 MongoDB, 连接指定的文档集合
func Collection(c *gin.Context) *mongo.Collection {
	// 从配置文件中读取连接配置
	uri := "mongodb://" + application.App.Mongo.Username + ":" + application.App.Mongo.Password + "@" + application.App.Mongo.Host + ":" + application.App.Mongo.Port + "/"

	// 连接 MongoDB 数据库
	client, err := mongo.Connect(c, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}

	// 连接配置文件指定的数据库和文档集合
	collection := client.Database(application.App.Mongo.Database).Collection(application.App.Mongo.Collection)

	return collection
}

// One 按 ObjectId 查询单条 Linux 命令
func One(c *gin.Context, commandId string) model.Command {
	// 获取集合连接
	collection := Collection(c)

	// 字符串 Id 转换为 ObjectId
	objectId, errHex := primitive.ObjectIDFromHex(commandId)
	if errHex != nil {
		log.Println(errHex)
	}

	// 按 ObjectId 查询数据
	result := collection.FindOne(c, bson.M{
		"_id": objectId,
	})

	// 结构体对象
	var command model.Command
	// 将查询结果赋值给结构体对象
	err := result.Decode(&command)
	if err != nil {
		log.Println(err)
	}

	return command
}

// OneByName 按 Linux 命令名称查询单条 Linux 命令
func OneByName(c *gin.Context, commandName string) model.Command {
	// 获取集合连接
	collection := Collection(c)

	// 按 Linux 命令名称查询数据
	result := collection.FindOne(c, bson.M{
		"command": commandName,
	})

	// 结构体对象
	var command model.Command
	// 将查询结果赋值给结构体对象
	err := result.Decode(&command)
	if err != nil {
		log.Println(err)
	}

	return command
}

// List 查询所有 Linux 命令
func List(c *gin.Context) []model.Command {
	// 获取集合连接
	collection := Collection(c)

	// 查询所有 Linux 命令
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 结构体数组
	var commandArray []model.Command
	// 返回值 cursor 相当于一个指针, 需要 Next() 遍历一个一个获取数据
	for cursor.Next(c) {
		command := model.Command{}
		err := cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		commandArray = append(commandArray, command)
	}

	return commandArray
}

// ListByName 查询所有 Linux 命令的名称
func ListByName(c *gin.Context) []string {
	// 获取集合连接
	collection := Collection(c)

	// 查询所有 Linux 命令
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 字符串数组
	var nameArray []string
	// 返回值 cursor 相当于一个指针, 需要 Next() 遍历一个一个获取数据
	for cursor.Next(c) {
		command := model.Command{}
		err := cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		nameArray = append(nameArray, command.Command)
	}

	return nameArray
}

// InsertOne 插入单条 Linux 命令
func InsertOne(c *gin.Context) (*mongo.InsertOneResult, string) {
	// 获取集合连接
	collection := Collection(c)

	// 结构体对象
	var command model.Command
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
	// 获取集合连接
	collection := Collection(c)

	// interface 数组
	var commandList []interface{}
	// 将请求体参数赋值到 interface 数组上
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

// UpdateOne 更新单条 Linux 命令
func UpdateOne(c *gin.Context) *mongo.UpdateResult {
	// 获取集合连接
	collection := Collection(c)

	// 结构体对象
	var command model.Command
	// 将请求体参数赋值到结构体对象上
	errBind := c.ShouldBind(&command)
	if errBind != nil {
		log.Println(errBind)
	}

	// 根据 ObjectId 更新单条 Linux 命令
	result, err := collection.UpdateOne(c, bson.M{"_id": command.Id}, bson.M{"$set": command})
	if err != nil {
		log.Println(err)
	}

	return result
}

// DeleteOne 删除单条 Linux 命令
func DeleteOne(c *gin.Context, commandId string) (*mongo.DeleteResult, primitive.ObjectID) {
	// 获取集合连接
	collection := Collection(c)

	// 字符串 Id 转换为 ObjectId
	objectId, errHex := primitive.ObjectIDFromHex(commandId)
	if errHex != nil {
		log.Println(errHex)
	}

	// 根据 ObjectId 删除单条 Linux 命令
	result, err := collection.DeleteOne(c, bson.M{"_id": objectId})
	if err != nil {
		log.Println(err)
	}

	return result, objectId
}

// DeleteMany 删除多条 Linux 命令
func DeleteMany(c *gin.Context) (*mongo.DeleteResult, []primitive.ObjectID) {
	// 获取集合连接
	collection := Collection(c)

	// 字符串 Id 数组
	var commandIds []string
	// 将请求体参数赋值到字符串 Id 数组上
	errBind := c.ShouldBind(&commandIds)
	if errBind != nil {
		log.Println(errBind)
	}

	// ObjectId 数组
	var objectIds []primitive.ObjectID
	// 遍历字符串 Id 数组, 转换为 ObjectId 数组
	for _, commandId := range commandIds {
		objectId, errHex := primitive.ObjectIDFromHex(commandId)
		if errHex != nil {
			log.Println(errHex)
		}
		objectIds = append(objectIds, objectId)
	}

	// 根据 ObjectId 数组删除多条 Linux 命令
	result, err := collection.DeleteMany(c, bson.M{"_id": bson.M{"$in": objectIds}})
	if err != nil {
		log.Println(err)
	}

	return result, objectIds
}

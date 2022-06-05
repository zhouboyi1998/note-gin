package repository

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
	uri := "mongodb://" +
		application.App.Mongodb.Username + ":" +
		application.App.Mongodb.Password + "@" +
		application.App.Mongodb.Host + ":" +
		application.App.Mongodb.Port + "/"

	// 连接 MongoDB 数据库
	client, err := mongo.Connect(c, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}

	// 连接配置文件指定的数据库和文档集合
	collection := client.Database(application.App.Mongodb.Database).Collection(application.App.Mongodb.Collection)

	return collection
}

// One 根据id查询命令
func One(c *gin.Context) model.Command {
	// 获取集合连接
	collection := Collection(c)

	// 获取命令id参数
	commandId := c.Param("commandId")
	// 转换成文档id
	objectId, errHex := primitive.ObjectIDFromHex(commandId)
	if errHex != nil {
		log.Println(errHex)
	}

	// 根据文档id查询命令
	result := collection.FindOne(c, bson.M{
		"_id": objectId,
	})

	// 将数据解码成命令对象
	var command model.Command
	err := result.Decode(&command)
	if err != nil {
		log.Println(err)
	}

	return command
}

// List 查询命令列表
func List(c *gin.Context) []model.Command {
	// 获取集合连接
	collection := Collection(c)

	// 查询命令列表
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 返回值数组
	var commandArray []model.Command
	// 使用 cursor 指针遍历数据
	for cursor.Next(c) {
		// 将数据解码成命令对象
		command := model.Command{}
		err := cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		// 添加到返回值数组中
		commandArray = append(commandArray, command)
	}

	return commandArray
}

// Insert 新增命令
func Insert(c *gin.Context) (*mongo.InsertOneResult, string) {
	// 获取集合连接
	collection := Collection(c)

	// 获取请求体参数
	var command model.Command
	errBind := c.ShouldBind(&command)
	if errBind != nil {
		log.Println(errBind)
	}
	// 生成文档id
	command.Id = primitive.NewObjectID()

	// 新增命令
	result, err := collection.InsertOne(c, command)
	if err != nil {
		log.Println(err)
	}

	return result, command.Command
}

// InsertBatch 批量新增命令
func InsertBatch(c *gin.Context) *mongo.InsertManyResult {
	// 获取集合连接
	collection := Collection(c)

	// 获取请求体参数
	var commandList []interface{}
	errBind := c.ShouldBind(&commandList)
	if errBind != nil {
		log.Println(errBind)
	}

	// 批量新增命令
	result, err := collection.InsertMany(c, commandList)
	if err != nil {
		log.Println(err)
	}

	return result
}

// Update 修改命令
func Update(c *gin.Context) *mongo.UpdateResult {
	// 获取集合连接
	collection := Collection(c)

	// 获取请求体参数
	var command model.Command
	errBind := c.ShouldBind(&command)
	if errBind != nil {
		log.Println(errBind)
	}

	// 根据命令id修改命令
	result, err := collection.UpdateByID(c, command.Id, bson.M{"$set": command})
	if err != nil {
		log.Println(err)
	}

	return result
}

// UpdateBatch 批量修改命令
func UpdateBatch(c *gin.Context) []*mongo.UpdateResult {
	// 获取集合连接
	collection := Collection(c)

	// 获取请求体参数
	var commandArray []model.Command
	errBind := c.ShouldBind(&commandArray)
	if errBind != nil {
		log.Println(errBind)
	}

	// 返回值数组
	var resultArray []*mongo.UpdateResult
	// 遍历需要修改的命令
	for _, command := range commandArray {
		// 根据命令id修改命令
		result, err := collection.UpdateByID(c, command.Id, bson.M{"$set": command})
		if err != nil {
			log.Println(err)
		}
		resultArray = append(resultArray, result)
	}

	return resultArray
}

// Delete 删除命令
func Delete(c *gin.Context) (*mongo.DeleteResult, primitive.ObjectID) {
	// 获取集合连接
	collection := Collection(c)

	// 获取命令id参数
	commandId := c.Param("commandId")
	// 转换成文档id
	objectId, errHex := primitive.ObjectIDFromHex(commandId)
	if errHex != nil {
		log.Println(errHex)
	}

	// 根据文档id删除命令
	result, err := collection.DeleteOne(c, bson.M{"_id": objectId})
	if err != nil {
		log.Println(err)
	}

	return result, objectId
}

// DeleteBatch 批量删除命令
func DeleteBatch(c *gin.Context) (*mongo.DeleteResult, []primitive.ObjectID) {
	// 获取集合连接
	collection := Collection(c)

	// 获取请求体参数
	var commandIds []string
	errBind := c.ShouldBind(&commandIds)
	if errBind != nil {
		log.Println(errBind)
	}

	// 遍历命令id数组, 转换为文档id
	var objectIds []primitive.ObjectID
	for _, commandId := range commandIds {
		objectId, errHex := primitive.ObjectIDFromHex(commandId)
		if errHex != nil {
			log.Println(errHex)
		}
		objectIds = append(objectIds, objectId)
	}

	// 根据文档id数组批量删除命令
	result, err := collection.DeleteMany(c, bson.M{"_id": bson.M{"$in": objectIds}})
	if err != nil {
		log.Println(err)
	}

	return result, objectIds
}

// Select 查询命令
func Select(c *gin.Context) model.Command {
	// 获取集合连接
	collection := Collection(c)

	// 获取命令名称参数
	commandName := c.Param("commandName")

	// 根据命令名称查询数据
	result := collection.FindOne(c, bson.M{
		"command": commandName,
	})

	// 将数据解码成命令对象
	var command model.Command
	err := result.Decode(&command)
	if err != nil {
		log.Println(err)
	}

	return command
}

// NameList 查询命令名称列表
func NameList(c *gin.Context) []string {
	// 获取集合连接
	collection := Collection(c)

	// 查询命令列表
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Println(err)
	}

	// 返回值数组
	var nameArray []string
	// 使用 cursor 指针遍历获取数据
	for cursor.Next(c) {
		// 将数据解码成命令对象
		command := model.Command{}
		err := cursor.Decode(&command)
		if err != nil {
			log.Println(err)
		}
		// 获取命令名称, 添加到返回值数组中
		nameArray = append(nameArray, command.Command)
	}

	return nameArray
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Command Linux 命令结构体
type Command struct {
	Id      primitive.ObjectID `bson:"_id"`
	Command string             `bson:"command"`
	Usage   string             `bson:"usage"`
	Params  []struct {
		Param       string `bson:"param"`
		Description string `bson:"description"`
	} `bson:"params"`
}

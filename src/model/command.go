package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Command 命令结构体
type Command struct {
	Id      primitive.ObjectID `bson:"_id"`
	Command string             `bson:"command"`
	Usage   string             `bson:"usage"`
	Params  []Param            `bson:"params"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Command Linux 命令结构体
type Command struct {
	Id      primitive.ObjectID `bson:"_id"`
	Command string             `json:"command"`
	Usage   string             `json:"usage"`
	Params  []struct {
		Param       string `json:"param"`
		Description string `json:"description"`
	} `json:"params"`
}

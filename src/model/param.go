package model

// Param Linux 命令参数结构体
type Param struct {
	Param       string `bson:"param"`
	Description string `bson:"description"`
}

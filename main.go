package main

import (
	"note-gin/src/application"
	"note-gin/src/route"
)

func main() {
	// 使用自定义路由规则
	r := route.Route()
	// 启动项目, 从配置文件中读取监听地址
	r.Run(application.App.Server.Host + ":" + application.App.Server.Port)
}

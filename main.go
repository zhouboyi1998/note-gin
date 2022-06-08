package main

import "note-gin/src/route"

func main() {
	// 使用自定义路由规则
	r := route.Route()
	// 监听 18091
	r.Run("127.0.0.1:18091")
}

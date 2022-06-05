package main

import "note-gin/src/route"

func main() {
	// 使用 Gin 默认路由引擎
	r := route.Route()
	// 监听 127.0.0.1:18091
	r.Run("127.0.0.1:18091")
}

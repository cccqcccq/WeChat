package main

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"wx/model"
	"wx/routers"
	"wx/server"
)

// 创建一个默认的路由引擎
var r = gin.Default()

func main() {
	/*
		设置为生产环境
		gin.SetMode(gin.ReleaseMode)
		gin.Default必须在这句话下面
	*/

	// 要直接在session中存储结构体需要注册结构体
	gob.Register(model.UserModel{})
	go server.ClientList.Start()
	// 注册路由
	router()
	// 运行http服务
	r.Run(":8080")
}

// 路由和返回数据
func router() {
	// 注册首页路由
	routers.IndexRouter(r)
	// 注册api路由
	routers.ApiRouter(r)
}

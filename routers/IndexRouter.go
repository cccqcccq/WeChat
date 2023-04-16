package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"wx/controller"
	"wx/middlewares"
)

func IndexRouter(r *gin.Engine) {
	// 解析文件
	r.LoadHTMLGlob("view/*")
	// 处理静态文件,比原生的好理解而且方便,第一个参数就是路由,第二个参数就是文件的真实路径,当前目录需要用 ./
	r.Static("/static", "./static")
	// 连接redis
	store, _ := redis.NewStore(10, "tcp", ":6379", "", []byte("cccq"))
	// 将session注册到全局中间件
	r.Use(sessions.Sessions("SessionId", store))

	// 设置页面小图标
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("static/images/wx.png")
	})

	// 首页
	IndexController := Controller.IndexController{}
	r.GET("/", middlewares.IsLogin, IndexController.Index)
	// 登录页面
	r.GET("/login", middlewares.AutoLogin, IndexController.Login)
	// 注册页面
	r.GET("/register", middlewares.AutoLogin, IndexController.Register)
	// 退出登录
	r.GET("/logout", middlewares.IsLogin, IndexController.Logout)
}

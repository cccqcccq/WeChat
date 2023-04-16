package routers

import (
	"github.com/gin-gonic/gin"
	Controller "wx/controller"
	"wx/middlewares"
)

func ApiRouter(r *gin.Engine) {
	api := r.Group("/api")

	ApiController := Controller.ApiController{}
	// webSocket服务器
	api.GET("/chatSocket", ApiController.Chat)
	// api
	api.POST("/login", middlewares.AutoLogin, ApiController.Login)
	api.POST("/register", ApiController.Register)
	api.POST("/MyFriends", middlewares.IsLogin, ApiController.MyFriends)
	api.POST("/getChatHistory", middlewares.IsLogin, ApiController.GetChatHistory)
}

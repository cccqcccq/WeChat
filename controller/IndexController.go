package Controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"wx/model"
	"wx/server"
)

type IndexController struct {
}

func (this *IndexController) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func (this *IndexController) Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func (this *IndexController) Register(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}

func (this *IndexController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	// 在ClientList中将客户端下线
	userModel := session.Get("login").(model.UserModel)
	server.ClientList.Clients[userModel.Username] = nil
	// 清空session
	session.Clear()
	// 清除cookie
	c.SetCookie("token", "", -1, "", "", false, false)
	c.SetCookie("SessionId", "", -1, "", "", false, false)
	c.Redirect(302, "/login")
}

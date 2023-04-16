package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("login") == nil {
		c.Redirect(302, "/login")
		// 停止代码往后执行,return无效
		c.Abort()
	}
}

func AutoLogin(c *gin.Context) {
	cookie, _ := c.Cookie("token")
	if cookie != "" {
		session := sessions.Default(c)
		if session.Get("login") != nil {
			c.Redirect(302, "/")
			// 停止代码往后执行,return无效
			c.Abort()
		}
	}
}

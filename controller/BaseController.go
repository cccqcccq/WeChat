package Controller

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (this *BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

func (this *BaseController) Error(c *gin.Context, error string) {
	c.JSON(200, gin.H{
		"code":  500,
		"error": error,
	})
}

func (this *BaseController) SetToken(marshal []byte) string {
	header := `{"type":"jwt","alg":"SHA256"}`
	// 调用属性自动创建一个 Encoding结构体
	base := base64.StdEncoding
	// base64编码
	header = base.EncodeToString([]byte(header))
	payload := base.EncodeToString(marshal)
	secret := "cccq"
	// cccq是秘钥
	temp := sha256.Sum256([]byte(header + "." + payload + secret))
	return header + "." + payload + "." + hex.EncodeToString(temp[:])
}

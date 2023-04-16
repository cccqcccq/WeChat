package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
)

// CheckToken 检查token
func CheckToken(token string) bool {
	// 分别获取三段token
	arr := strings.Split(token, ".")
	header := arr[0]
	payload := arr[1]
	secret := "cccq"
	// 加密
	temp := sha256.Sum256([]byte(header + "." + payload + secret))
	// 三段拼接,判断是否一致
	rightToken := header + "." + payload + "." + hex.EncodeToString(temp[:])
	if token == rightToken {
		return true
	}
	return false
}

// TokenGetValue 从token中获取字段值
/*func TokenGetValue(token, key string) string {
	base := base64.StdEncoding
	data := make(map[string]interface{}, 0)
	// 获取token中的数据
	arr := strings.Split(token, ".")
	payload := arr[1]
	// 创建base解码获取json
	temp, _ := base.DecodeString(payload)
	json.Unmarshal(temp, &data)
	return data[key].(string)
}*/

// TokenGetID 从token中获取ID
func TokenGetID(token string) int {
	base := base64.StdEncoding
	data := make(map[string]interface{}, 0)
	// 获取token中的数据
	arr := strings.Split(token, ".")
	payload := arr[1]
	// 创建base解码获取json
	temp, _ := base.DecodeString(payload)
	json.Unmarshal(temp, &data)
	return int(data["id"].(float64))
}

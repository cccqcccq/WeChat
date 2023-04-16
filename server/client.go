package server

import (
	"encoding/json"
	"wx/utils"
)

func (this *Client) Read() {
	defer this.Conn.Close()
	for {
		//this.Conn.PongHandler()
		//读取ws中的数据
		m, message, err := this.Conn.ReadMessage()
		if err != nil {
			return
		}
		// 获取发送来的token
		data := make(map[string]interface{}, 0)
		json.Unmarshal(message, &data)
		token := data["token"]
		// 检查token
		if !utils.CheckToken(token.(string)) {
			sendMessage := SendMessage{Code: 500, Content: "token错误,请重新登录"}
			msg, _ := json.Marshal(sendMessage)
			this.Conn.WriteMessage(m, msg)
		} else {
			// 将消息存入列表
			ChatList.Message <- data
		}
	}
}

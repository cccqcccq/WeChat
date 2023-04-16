package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"strconv"
	"wx/utils"
)

func (this *ClientListStruct) Start() {
	redisConn := utils.Pool.Get()
	for {
		select {
		// 处理用户连接
		case client := <-ClientList.Register:
			// 将连接存入列表
			ClientList.Clients[client.ID] = client
		// 处理消息列表
		case chat := <-ChatList.Message:
			switch chat["type"] {
			case "text":
				// 发送用户和收消息用户的ID
				userId := utils.TokenGetID(chat["token"].(string))
				chatId := int(chat["chat_id"].(float64))
				key := ""
				if userId > chatId {
					key = strconv.Itoa(chatId) + "-" + strconv.Itoa(userId)
				} else {
					key = strconv.Itoa(userId) + "-" + strconv.Itoa(chatId)
				}
				// 将消息发送给对应的client
				client := ClientList.Clients[chat["send_id"].(string)]
				redisConn.Do("rpush", key, strconv.Itoa(userId)+":"+chat["content"].(string))
				if client != nil {
					sendMessage := SendMessage{Code: 200, UserId: utils.TokenGetID(chat["token"].(string)), Content: chat["content"].(string)}
					msg, _ := json.Marshal(sendMessage)
					client.Conn.WriteMessage(websocket.TextMessage, msg)
				}
			}
		}
	}
}

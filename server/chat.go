package server

import (
	"github.com/gorilla/websocket"
)

// 创建所需结构体

// Client 用户类
type Client struct {
	ID   string // 用户的username
	Conn *websocket.Conn
}

// ClientListStruct 用户列表
type ClientListStruct struct {
	Clients  map[string]*Client
	Register chan *Client
}

// ChatListStruct 消息列表
type ChatListStruct struct {
	Message chan map[string]interface{}
}

// SendMessage 消息
type SendMessage struct {
	Code    int    `json:"code"`
	UserId  int    `json:"userId"` // 发送消息的用户ID
	Content string `json:"content"`
}

// 实例化

// ClientList 用于存入客户端
var ClientList = ClientListStruct{
	Clients:  make(map[string]*Client),
	Register: make(chan *Client),
}

// ChatList 用于接收用户发送的消息
var ChatList = ChatListStruct{
	Message: make(chan map[string]interface{}),
}

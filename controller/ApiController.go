package Controller

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"wx/model"
	"wx/server"
	"wx/utils"
)

type ApiController struct {
	BaseController
}

func (this *ApiController) Login(c *gin.Context) {
	// 调用方法查看账号密码是否存在
	username := c.PostForm("username")
	password := c.PostForm("password")
	userModel := model.UserModel{Username: username, Password: password}
	user := userModel.UserLogin()
	if user == nil {
		this.Error(c, "账号或密码错误")
		return
	}
	session := sessions.Default(c)
	// 要直接在session中存储结构体需要注册结构体,main方法中注册了
	session.Set("login", user)
	session.Save()

	// 返回数据
	marshal, _ := json.Marshal(user)
	// 设置为token
	this.Success(c, this.SetToken(marshal))
}

func (this *ApiController) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	region := c.PostForm("region")
	if username == "" || password == "" || nickname == "" || region == "" {
		this.Error(c, "请确保每一项不为空")
		return
	}
	if len(password) < 6 {
		this.Error(c, "密码最少六位数")
		return
	}
	userModel := model.UserModel{Username: username, Password: password, Name: nickname, Area: region}
	err := userModel.UserRegister()
	if err != nil {
		this.Error(c, err.Error())
		return
	}
	this.Success(c, "注册成功,跳转到登录页面")
}

// Chat 用户进入首页时会调用这个方法形成一个连接
func (this *ApiController) Chat(c *gin.Context) {
	// 从session中获取用户信息
	session := sessions.Default(c)
	data := session.Get("login").(model.UserModel)
	// 创建连接
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// 创建一个用户实例
	client := &server.Client{
		ID:   data.Username,
		Conn: conn,
	}
	// 将用户加入列表
	server.ClientList.Register <- client
	// 开启读取这个用户的消息的协程
	go client.Read()
	//	go client.Write()
}

// GetChatHistory 获取聊天记录
func (this *ApiController) GetChatHistory(c *gin.Context) {
	page, _ := strconv.Atoi(c.PostForm("page"))
	session := sessions.Default(c)
	redisConn := utils.Pool.Get()
	// 从session中获取用户id
	userId := session.Get("login").(model.UserModel).Id
	chatId, _ := strconv.Atoi(c.PostForm("chatId"))
	// redis key的存储规则是小的id-大的id
	key := ""
	if userId > chatId {
		key = strconv.Itoa(chatId) + "-" + strconv.Itoa(userId)
	} else {
		key = strconv.Itoa(userId) + "-" + strconv.Itoa(chatId)
	}
	data := make([]string, 0)
	// 获取数据长度
	length, _ := redis.Int(redisConn.Do("llen", key))
	// 数据起始和截止
	start := length - ((page) * 20)
	if start < 0 {
		start = 0
	}
	stop := length - ((page - 1) * 20) - 1
	// 获取数据并返回,每次只获取20条,从后往前获取,也就是获取最新的聊天记录,redis获取规则,从第几条到第几条,与mysql不同
	values, _ := redis.Values(redisConn.Do("lrange", key, start, stop))
	// stop小于0代表数据已经全部返回过了,直接返回空数据
	if stop < 0 {
		this.Success(c, make([]string, 0))
		return
	}
	for _, v := range values {
		res, _ := redis.String(v, nil)
		data = append(data, res)
	}
	this.Success(c, data)
}

// MyFriends 获取用户好友列表
func (this *ApiController) MyFriends(c *gin.Context) {
	session := sessions.Default(c)
	// 获取好友信息
	friendsModel := model.FriendsModel{}
	// 从session获取用户id传入
	str, err := json.Marshal(friendsModel.GetFriends(session.Get("login").(model.UserModel).Id))
	if err != nil {
		return
	}
	this.Success(c, string(str))
}

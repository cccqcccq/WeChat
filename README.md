## 前言
感谢Delusion提供的页面支持,仿照pc端微信做的web页面,项目基于gin框架搭建,目前仅实现基础功能,未来可能会完善项目

### 使用的开源项目
| 开源项目                                                 | 地址                    |
| ---------------------------------------------------- |-----------------------|
| gin        | https://github.com/gin-gonic/gin |
| websocket  | https://github.com/gorilla/websocket |
| mysql      | https://github.com/go-sql-driver/mysql |
| redis      | https://github.com/gomodule/redigo/redis |
| session    | https://github.com/gin-contrib/sessions |

## 项目结构
```
├─controller -- 一控制器
│  ├─ApiController.go  -- api都在这个控制器返回
│  ├─BaseController.go  -- 实现共同返回json的方法
│  ├─IndexController.go  -- 用于返回页面
├─middlewares -- 中间件
│  ├─Login.go  -- 自动登录和判断是否登录
├─model -- 一数据库交互
│  ├─FriendsModel.go  -- 对应数据库中的friends表
│  ├─UserModel.go  -- 对应数据库中的user表
├─routers -- 一路由控制
│  ├─ApiRouter.go  -- api路由
│  ├─IndexRouter.go  -- 页面路由
├─server -- 一webscoket功能
│  ├─chat.go  -- 定义结构体
│  ├─client.go  -- 读取客户端发送的信息
│  ├─list.go  -- 处理客户端发送的信息
├─static -- 一静态资源
│  ├─css
│  ├─fonts
│  ├─images
│  ├─js
├─utils -- 一工具
│  ├─db.go  -- mysql连接
│  ├─redis.go  -- redis连接
│  ├─token.go  -- 处理token
├─view -- 一html视图
│  ├─index.html  -- 首页
│  ├─login.html  -- 登录页
│  ├─register.html  -- 注册页
```
## 项目运行
启动mysql,导入sql文件,在db.go中可以设置数据库账号密码

启动redis,端口号默认6379

运行main.go

项目地址:http://127.0.0.1:8080

两个账号密码均为123456

## 展望未来
* 好友/群聊
  * 添加/删除好友
  * 创建/退出群聊
  * 搜索好友
* 发送表情/文件
* 朋友圈
* 设置头像
* 聊天记录持久化
* 密码加盐

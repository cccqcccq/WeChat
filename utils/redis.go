package utils

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// Pool 定义一个全局的连接池
var Pool *redis.Pool

// go提供的初始化函数
func init() {
	// redis连接池,提前初始化一定数量的连接放入连接池,在需要用时从连接池内去吃,节省了每次需要时连接redis的时间
	Pool = &redis.Pool{
		MaxIdle:     8,                // 最大空闲连接数,初始化8个,不够时会自动创建连接放入连接池
		MaxActive:   0,                // 表示和数据库的最大连接数,0表示没限制
		IdleTimeout: time.Second * 10, // 最大空闲时间,如果连接池的的连接闲置时长超过该值就会关闭连接,0表示不关闭,这里表示的是10秒
		// 初始化连接redis的方法
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

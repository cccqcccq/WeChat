package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB // 全局变量连接数据库
	err error   // 数据库的错误
)

func init() {
	Db, err = sql.Open("mysql", "root:root@/wx")
	if err != nil {
		return
	}
}

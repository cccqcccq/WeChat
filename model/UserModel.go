package model

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"wx/utils"
)

type UserModel struct {
	Id       int    `json:"id"`
	Username string `json:"username,omitempty"`
	Password string `json:"-"` // json序列化时隐藏这个字段
	Name     string `json:"name,omitempty"`
	Area     string `json:"area,omitempty"`
	Image    string `json:"image,omitempty"`
}

// UserLogin 账号是否存在,存在返回用户信息
func (this *UserModel) UserLogin() *UserModel {
	sql := "select * from user where username=? and password=?"
	temp := sha256.Sum256([]byte(this.Password))
	// 密码加密
	this.Password = hex.EncodeToString(temp[:])
	// 只查询一条数据
	row := utils.Db.QueryRow(sql, this.Username, this.Password)
	// 将用户信息存入结构体返回
	err := row.Scan(&this.Id, &this.Username, &this.Password, &this.Name, &this.Area, &this.Image)
	if err != nil {
		return nil
	}
	return this
}

func (this *UserModel) UserRegister() error {
	sql := "select id from user where username=?"
	row := utils.Db.QueryRow(sql, this.Username)
	row.Scan(&this.Id)
	if this.Id != 0 {
		return errors.New("账号已存在")
	}
	temp := sha256.Sum256([]byte(this.Password))
	// 密码加密
	this.Password = hex.EncodeToString(temp[:])
	sql = "insert into user(username,password,name,area) values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, this.Username, this.Password, this.Name, this.Area)
	if err != nil {
		fmt.Println(err)
		return errors.New("注册失败,请重试")
	}
	return nil
}

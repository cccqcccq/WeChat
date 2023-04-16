package model

import (
	"fmt"
	"wx/utils"
)

type FriendsModel struct {
	Id       int `json:"id"`
	UserId   int `json:"user_id"`
	FriendId int `json:"friend_id"`
}

func (this *FriendsModel) GetFriends(id int) []UserModel {
	/*
		如果选中了f表的id但是不想要这个字段,需要一个没用的变量用来接收scan的返回,不能用nil
		sql := "select f.id,u.* from friends f join user u on u.id=f.friend_id where f.user_id=?"
		temp := ""
		query.Scan(&temp2,&temp.Id, &temp.Username, &temp.Password, &temp.Name, &temp.Area, &temp.Image)
	*/

	// 获取所有好友的信息
	sql := "select u.* from friends f join user u on u.id=f.friend_id where f.user_id=?"
	query, err := utils.Db.Query(sql, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data := make([]UserModel, 0)
	// 将数据存入user结构体,再存入切片中
	temp := UserModel{}
	for query.Next() {
		query.Scan(&temp.Id, &temp.Username, &temp.Password, &temp.Name, &temp.Area, &temp.Image)
		data = append(data, temp)
	}
	return data
}

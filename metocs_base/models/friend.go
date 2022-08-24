package models

import (
	"gorm.io/gorm"
	"metocs/base/database"
)

type Friend struct {
	gorm.Model

	UserId   string `json:"UserId"`
	FriendId string `json:"friendId" binding:"required"`
	MarkName string `json:"markName"`
	Label    string `json:"label"`
}

func (friend *Friend) AddFriend() {
	tx := database.DB.Create(friend)
	if tx.Error != nil {
		panic("请勿重复添加！")
	}
}

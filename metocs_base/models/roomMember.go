package models

import (
	"gorm.io/gorm"
	"metocs/base/database"
)

type RoomMember struct {
	gorm.Model
	RoomId     uint
	Member     string
	Status     uint
	MemberType uint
}

func CreateRoom(roomId uint, member string) *RoomMember {

	roomMember := &RoomMember{
		RoomId: roomId,
		Member: member,
	}
	database.DB.Create(roomMember)
	return roomMember
}

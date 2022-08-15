package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	RoomId     uint
	UserId     string
	Message    string
	Status     int8
	MemberType int8
}

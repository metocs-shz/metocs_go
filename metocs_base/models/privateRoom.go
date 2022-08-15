package models

import "gorm.io/gorm"

type (
	PrivateRoom struct {
		AbstractRoom
	}

	AbstractRoom struct {
		gorm.Model
		Name    string          `json:"name" gorm:"column:name"`
		UserId  string          `json:"userId"`
		Members map[string]User `json:"members" gorm:"-"`
	}
)

func (room AbstractRoom) GetRoomId() uint {
	return room.ID
}

func (room AbstractRoom) GetRoomName() string {
	return room.Name
}

func (room AbstractRoom) GetRoomMembers() map[string]User {
	return room.Members
}

func (room AbstractRoom) AddMember(members ...User) {
	if room.Members == nil {
		room.Members = make(map[string]User)
	}
	for _, member := range members {
		room.Members[member.ID] = member
	}
}

package models

import "gorm.io/gorm"

type (
	PrivateRoom struct {
		AbstractRoom
	}

	AbstractRoom struct {
		gorm.Model
		Name    string            `json:"name" gorm:"column:name"`
		Members map[string]Member `gorm:"-"`
	}
)

func (room AbstractRoom) GetRoomId() uint {
	return room.ID
}

func (room AbstractRoom) GetRoomName() string {
	return room.Name
}

func (room AbstractRoom) GetRoomMembers() map[string]Member {
	return room.Members
}

func (room AbstractRoom) AddMember(members ...Member) {
	if room.Members == nil {
		room.Members = make(map[string]Member)
	}
	for _, member := range members {
		room.Members[member.Id] = member
	}
}

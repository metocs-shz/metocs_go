package models

//Room 房间接口
type Room interface {

	// GetRoomId 获取房间id
	GetRoomId() uint

	// GetRoomName 获取房间名称
	GetRoomName() string

	// GetRoomMembers 湖区房间内成员列表
	GetRoomMembers() map[uint]User

	// AddMember 向房间中添加成员
	AddMember(member ...User)
}

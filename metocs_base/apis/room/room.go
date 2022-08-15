package room

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"metocs_go/database"
	"metocs_go/models"
	"metocs_go/utils"
)

//RoomDataParam 房间参数
type RoomDataParam struct {
	Name    string   `json:"name"`
	Member  string   `json:"member"`
	Members []string `json:"members"`
}

// CreateRoom Create 创建一个房间
func CreateRoom(c *gin.Context) {

	param := &RoomDataParam{}
	err := c.ShouldBindBodyWith(param, binding.JSON)
	if err != nil {
		log.Println(err.Error())
		utils.Fail(c, 500, "数据绑定失败！")
	}

	id := utils.GetClaim(c).ID

	// 创建房间
	var room = &models.PrivateRoom{}
	room.Name = param.Name
	room.UserId = id
	database.DB.Create(room)

	var roomMember = &models.RoomMember{}
	roomMember.RoomId = room.ID
	roomMember.Member = id

	utils.Success(c, room)
}

package room

import (
	"github.com/gin-gonic/gin"
	"metocs/base/apis/room"
)

func RouterGroup(router *gin.RouterGroup) {

	group := router.Group("/room")
	{
		group.POST("/create", room.CreateRoom)

	}

}

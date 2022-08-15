package router

import (
	"github.com/gin-gonic/gin"
	"metocs_go/apis"
)

func RoomRouterGroup(router *gin.RouterGroup) {

	group := router.Group("/room")
	{
		group.GET("/create", apis.Create)

	}

}

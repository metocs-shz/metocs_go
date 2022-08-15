package auth

import (
	"github.com/gin-gonic/gin"
	"metocs_go/apis/auth"
)

func RouterGroup(router *gin.RouterGroup) {

	group := router.Group("/auth")
	{
		group.POST("/login", auth.Login)
		//	group.GET("/friends", friend.GetFriends)

	}

}

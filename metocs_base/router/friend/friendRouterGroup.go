package friend

import (
	"github.com/gin-gonic/gin"
	"metocs_go/apis/friend"
)

func RouterGroup(router *gin.RouterGroup) {

	group := router.Group("/friend")
	{
		group.POST("/add", friend.AddFriend)
		group.GET("/friends", friend.GetFriends)

	}

}

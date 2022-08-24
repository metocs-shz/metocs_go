package user

import (
	"github.com/gin-gonic/gin"
	"metocs/base/apis/user"
)

func RouterGroup(router *gin.RouterGroup) {

	group := router.Group("/user")
	{
		group.POST("/create", user.CreateUser)
		group.GET("/getUserByID", user.GetUserByID)
	}

}

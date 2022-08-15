package message

import (
	"github.com/gin-gonic/gin"
	"metocs_go/apis/message"
)

func RouterGroup(router *gin.RouterGroup) {

	group := router.Group("/message")
	{
		group.POST("/send", message.SendMessage)
	}

}

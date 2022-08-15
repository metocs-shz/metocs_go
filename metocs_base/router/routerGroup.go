package router

import (
	"github.com/gin-gonic/gin"
	"metocs_go/config"
	"metocs_go/middleware"
	"metocs_go/router/auth"
	"metocs_go/router/friend"
	"metocs_go/router/message"
	"metocs_go/router/room"
	"metocs_go/router/user"
)

func GroupInit(engine *gin.Engine) {
	app := config.App
	group := engine.Group(app.Server.Path)
	{
		auth.RouterGroup(group)

		group.Use(middleware.Authorize())
		user.RouterGroup(group)
		friend.RouterGroup(group)
		room.RouterGroup(group)
		message.RouterGroup(group)
	}
}

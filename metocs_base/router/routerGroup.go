package router

import (
	"github.com/gin-gonic/gin"
	"metocs/base/config"
	"metocs/base/middleware"
	"metocs/base/router/auth"
	"metocs/base/router/friend"
	"metocs/base/router/message"
	"metocs/base/router/room"
	"metocs/base/router/user"
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

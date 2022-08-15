package router

import (
	"github.com/gin-gonic/gin"
	"metocs_go/config"
)

func GroupInit(engine *gin.Engine) {
	app := config.App
	group := engine.Group(app.Server.Path)
	{
		RoomRouterGroup(group)
		UserRouterGroup(group)
	}
}

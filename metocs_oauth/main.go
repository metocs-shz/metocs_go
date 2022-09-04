package main

import (
	"github.com/gin-gonic/gin"
	"metocs/common/middlerware"
	"metocs/common/run"
	"metocs/oauth/manager"
	"metocs/oauth/router"
)

func main() {
	run.ApplicationInit()
	manager.OauthInit()
	engine := gin.Default()
	engine.NoMethod(middlerware.Recover405)
	engine.NoRoute(middlerware.Recover404)
	router.BaseGroup(engine)
	run.Run(engine)
}

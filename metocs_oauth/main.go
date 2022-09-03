package main

import (
	"github.com/gin-gonic/gin"
	"metocs/common/run"
	"metocs/oauth/manager"
	"metocs/oauth/router"
)

func main() {
	run.ApplicationInit()
	manager.OauthInit()
	engine := gin.Default()
	router.BaseGroup(engine)
	run.Run(engine)
}

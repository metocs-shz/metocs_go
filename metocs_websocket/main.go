package main

import (
	"github.com/gin-gonic/gin"
	"metocs/common/middlerware"
	"metocs/common/run"
	"metocs/websocket/router"
)

func main() {
	run.ApplicationInit()
	engine := gin.Default()
	engine.NoMethod(middlerware.Recover405)
	engine.NoRoute(middlerware.Recover404)
	router.BaseGroup(engine)
	run.Run(engine)
}

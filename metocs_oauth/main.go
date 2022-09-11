package main

import (
	"metocs/common/base"
	"metocs/common/run"
	"metocs/oauth/manager"
	"metocs/oauth/router"
)

func main() {
	run.ApplicationInit()
	manager.OauthInit()
	engine := base.InitRouter()
	router.BaseGroup(engine)
	run.Run(engine)
}

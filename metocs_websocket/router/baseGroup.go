package router

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"metocs/common/application"
	"metocs/common/middlerware"
)

func BaseGroup(engine *gin.Engine) {
	server := application.Application.Server
	router := engine.Group(server.Path)

	engine.Group(server.Path + "/open")
	{

	}

	router.Use(middlerware.Recover).Use(ginserver.HandleTokenVerify(middlerware.MetocsConfig))
	{

	}
}

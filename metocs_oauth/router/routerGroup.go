package router

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"metocs/common/application"
	"metocs/oauth/apis"
)

func BaseGroup(engine *gin.Engine) {
	server := application.Application.Server

	router := engine.Group(server.Path)

	{
		ClientGroup(router)
		UserGroup(router)
		AuthGroup(router)
	}

	open := engine.Group(server.Path + "/open")
	{
		open.GET("/authorize", apis.Authorize)
		open.POST("/tokenClient", apis.ClientToken)
		open.GET("/token", ginserver.HandleTokenRequest)
		open.POST("/login", apis.Login)

	}
}

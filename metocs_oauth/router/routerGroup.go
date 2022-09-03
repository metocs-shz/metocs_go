package router

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"metocs/common/application"
	"metocs/oauth/apis/oauth"
	"metocs/oauth/models"
)

func BaseGroup(engine *gin.Engine) {
	server := application.Application.Server
	group := engine.Group(server.Path)
	{

		group.GET("/authorize", oauth.Authorize)
		group.GET("/token", ginserver.HandleTokenRequest)

		//以下接口需要权限
		group.Use(ginserver.HandleTokenVerify(models.MetocsConfig))

		group.Group("/client", ClientGroup)
		group.Group("/user", UserGroup)
		group.Group("/auth", AuthGroup)
	}
}

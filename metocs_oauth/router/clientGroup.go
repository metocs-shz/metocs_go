package router

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"metocs/oauth/apis"
	"metocs/oauth/models"
)

// ClientGroup 客户端管理分组
func ClientGroup(group *gin.RouterGroup) {

	engine := group.Group("/client").Use(ginserver.HandleTokenVerify(models.MetocsConfig))
	{
		engine.POST("/create", apis.CreateClient)
		engine.DELETE("/deleted", apis.DeletedClient)
		engine.GET("/page", apis.ClientPage)
	}

}

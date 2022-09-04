package router

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"metocs/oauth/apis"
	"metocs/oauth/models"
)

// UserGroup 用户管理分组
func UserGroup(group *gin.RouterGroup) {

	engine := group.Group("/user").Use(ginserver.HandleTokenVerify(models.MetocsConfig))
	{
		engine.POST("/create", apis.CreateUser)
		engine.DELETE("/deleted", apis.DeletedUser)
		//engine.PUT("/update", apis.Create)
		//engine.GET("/one", apis.Create)
		//	engine.GET("/page", apis.Create)
	}
}

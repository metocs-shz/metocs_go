package router

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"metocs/common/middlerware"
	"metocs/oauth/apis"
)

// UserGroup 用户管理分组
func UserGroup(group *gin.RouterGroup) {

	engine := group.Group("/user").Use(ginserver.HandleTokenVerify(middlerware.MetocsConfig))
	{
		engine.POST("/create", apis.CreateUser)
		engine.DELETE("/deleted", apis.DeletedUser)
		engine.PUT("/update", apis.UpdateUser)
		engine.GET("/one", apis.OneUser)
		engine.GET("/page", apis.UserPage)
	}
}

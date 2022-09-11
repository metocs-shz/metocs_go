package apis

import (
	"github.com/gin-gonic/gin"
	"metocs/common/application"
	"metocs/oauth/manager"
	"net/http"
)

func Authorize(context *gin.Context) {
	err := manager.OauthServer.HandleAuthorizeRequest(context.Writer, context.Request)
	if err != nil {
		context.JSON(http.StatusAccepted, gin.H{"code": "2", "message": "授权失败", "err": err.Error()})
	}
}

func AuthorizeOauth(context *gin.Context) {
	err := manager.OauthServer.HandleAuthorizeRequest(context.Writer, context.Request)
	if err != nil {
		context.JSON(http.StatusAccepted, gin.H{"code": "2", "message": "授权失败", "err": err.Error()})
	}
}

func ClientToken(context *gin.Context) {
	err := manager.OauthServer.HandleTokenRequest(context.Writer, context.Request)
	if err != nil {
		context.JSON(http.StatusAccepted, gin.H{"code": "2", "message": "授权失败", "err": err.Error()})
	}
}

type LoginModule struct {
	Phone    string
	Password string
}

func Login(context *gin.Context) {
	loginModule := &LoginModule{}
	err := context.ShouldBindJSON(loginModule)
	if err != nil {
		return
	}

	auth := application.Application.Auth
	m := make(map[string][]string)
	m["username"] = []string{loginModule.Phone}
	m["password"] = []string{loginModule.Password}
	m["grant_type"] = []string{auth.GrantType}
	m["client_id"] = []string{auth.ClientId}
	m["client_secret"] = []string{auth.ClientSecret}
	context.Request.Form = m

	err = manager.OauthServer.HandleTokenRequest(context.Writer, context.Request)
	if err != nil {
		context.JSON(http.StatusAccepted, gin.H{"code": "2", "message": "授权失败", "err": err.Error()})
	}
}

package apis

import (
	"github.com/gin-gonic/gin"
	"metocs/oauth/manager"
	"net/http"
)

func Authorize(context *gin.Context) {
	err := manager.OauthServer.HandleAuthorizeRequest(context.Writer, context.Request)
	if err != nil {
		context.JSON(http.StatusAccepted, gin.H{"code": "2", "message": "授权失败", "err": err.Error()})
	}
}

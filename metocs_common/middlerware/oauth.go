package middlerware

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"net/http"
)

var MetocsConfig = ginserver.Config{
	ErrorHandleFunc: ErrorHandleFunc,
	TokenKey:        "metocs_token",
	Skipper: func(_ *gin.Context) bool {
		return false
	},
}

func ErrorHandleFunc(c *gin.Context, error error) {
	c.JSON(http.StatusOK, gin.H{"code": 401, "message": error.Error()})
	c.Abort()
	return
}

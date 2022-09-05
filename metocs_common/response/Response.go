package response

import (
	"github.com/gin-gonic/gin"
	"metocs/common/database"
	"net/http"
)

const (
	SUCCESS_CODE         = 0
	SERVER_ERROR         = 1
	BAD_REQUEST_PARAM    = 2
	METHOD_IS_NOT_ALLOWD = 3
	HANDLER_NOT_Found    = 4
)

type page struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": SUCCESS_CODE, "message": "成功！"})
}

func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": SUCCESS_CODE, "message": "成功！", "data": data})
}

func SuccessPage(c *gin.Context, data interface{}) {
	var count int64
	database.DB.Find(data).Count(&count)
	page := new(page)
	page.List = data
	page.Total = count
	c.JSON(http.StatusOK, gin.H{"code": SUCCESS_CODE, "message": "成功！", "data": page})
}

func BadRequestError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": BAD_REQUEST_PARAM, "message": "请求参数错误！"})
}

func ServerError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": SERVER_ERROR, "message": "哎呀！服务器出现错误了！"})
}

func ServerErrorMessage(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"code": SERVER_ERROR, "message": "哎呀！服务器出现错误了！", "error": err.Error()})
}

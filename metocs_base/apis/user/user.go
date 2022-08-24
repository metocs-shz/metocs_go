package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"metocs/base/models"
	"metocs/base/utils"
)

// CreateUser 创建一个用户
func CreateUser(c *gin.Context) {
	user := &models.User{}
	err := c.BindJSON(user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	models.CreateUser(user)
	utils.SuccessNoMessage(c)
}

// GetUserByID 通过Id 获取永固信息
func GetUserByID(c *gin.Context) {
	param := c.Query("id")
	user := &models.User{}
	user.GetUserByID(param)
	utils.Success(c, user)
}

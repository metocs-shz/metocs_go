package apis

import (
	"github.com/gin-gonic/gin"
	"metocs/common/authTool"
	"metocs/common/database"
	"metocs/common/response"
	"metocs/oauth/models"
)

func CreateUser(c *gin.Context) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		response.BadRequestError(c)
	}
	user.PassWord = authTool.PassWordEncoder(user.PassWord)
	tx := database.DB.Create(user)
	if tx.Error != nil {
		response.ServerErrorMessage(c, tx.Error)
	}
	response.Success(c)
}

func DeletedUser(c *gin.Context) {
	id, b := c.GetQuery("id")
	if !b {
		response.BadRequestError(c)
	}
	database.DB.Delete("id = ?", id)
	response.Success(c)
}

package apis

import (
	"github.com/gin-gonic/gin"
	"metocs/common/authTool"
	"metocs/common/database"
	"metocs/common/response"
	"metocs/common/utils"
	"metocs/oauth/models"
)

func CreateUser(c *gin.Context) {
	user := &models.User{}
	user.ID = utils.GetID()
	err := c.ShouldBindJSON(user)
	if err != nil {
		response.BadRequestError(c)
		return
	}
	user.PassWord = authTool.PassWordEncoder(user.PassWord)
	tx := database.DB.Create(user)
	if tx.Error != nil {
		response.ServerErrorMessage(c, tx.Error)
		return
	}
	response.Success(c)
}

func DeletedUser(c *gin.Context) {
	id, b := c.GetQuery("id")
	if !b {
		response.BadRequestError(c)
		return
	}
	database.DB.Delete("id = ?", id)
	response.Success(c)
}

func UpdateUser(c *gin.Context) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil || user.ID == "" {
		response.BadRequestError(c)
		return
	}
	if user.PassWord != "" {
		user.PassWord = authTool.PassWordEncoder(user.PassWord)
	}
	tx := database.DB.Updates(user)
	if tx.Error != nil {
		response.ServerErrorMessage(c, tx.Error)
		return
	}
	response.Success(c)
}

func OneUser(c *gin.Context) {
	id, b := c.GetQuery("id")
	if !b {
		response.BadRequestError(c)
	}
	user := &models.User{}
	tx := database.DB.Where("id = ?", id).First(user)
	if tx.Error != nil {
		response.ServerErrorMessage(c, tx.Error)
		return
	}
	response.Success(c)
}

func UserPage(c *gin.Context) {
	users := &[]models.User{}
	pageSize := c.GetInt("pageSize")
	pageNum := c.GetInt("pageNum")
	query, b := c.GetQuery("name")

	db := database.DB
	if b {
		db.Where("name like '%?%' ", query)
	}

	tx := db.Limit(pageSize).Offset(utils.GetPage(pageNum, pageSize)).Order("id desc").Find(users)
	if tx.Error != nil {
		response.ServerErrorMessage(c, tx.Error)
		return
	}
	response.SuccessPage(c, users)

}

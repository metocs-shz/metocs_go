package apis

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"metocs/common/database"
	"metocs/common/response"
	"metocs/common/utils"
	"strings"
	"time"
)

type Oauth2Clients struct {
	ID        string `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Secret    string `json:"secret"`
	Domain    string `json:"domain"`
	Data      string `json:"data"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateClient(c *gin.Context) {
	client := &Oauth2Clients{}
	c.ShouldBindJSON(client)
	if client.Data == "" {
		client.Data = "{}"
	}

	client.ID = strings.ToUpper(utils.Md5(15))
	client.Secret = strings.ToUpper(utils.Md5(20))

	err := database.DB.Create(client).Error
	if err != nil {
		response.ServerErrorMessage(c, err)
	}
	response.SuccessData(c, client)
}

func DeletedClient(c *gin.Context) {
	id, b := c.GetQuery("id")
	if !b {
		response.BadRequestError(c)
	}
	database.DB.WithContext(c).Table("oauth2_clients").Delete("id = ?", id)
	response.Success(c)
}

func ClientPage(c *gin.Context) {
	clients := &[]Oauth2Clients{}
	pageSize := c.GetInt("pageSize")
	pageNum := c.GetInt("pageNum")
	query, b := c.GetQuery("name")
	var str = ""
	var err error
	if !b {
		str = "JSON_EXTRACT(content,'$.name')=(?)"
		err = database.DB.Where(str, query).Limit(pageSize).Offset(utils.GetPage(pageNum, pageSize)).Order("created_at desc").Find(clients).Error
	} else {
		err = database.DB.Limit(pageSize).Offset(utils.GetPage(pageNum, pageSize)).Order("created_at desc").Find(clients).Error
	}
	if err != nil {
		response.ServerError(c)
	}

	response.SuccessData(c, clients)
}

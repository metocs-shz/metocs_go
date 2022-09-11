package apis

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

}

func DeletedClient(c *gin.Context) {

}

func ClientPage(c *gin.Context) {

}

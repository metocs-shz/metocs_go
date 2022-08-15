package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"metocs_go/database"
	"metocs_go/models"
	"net/http"
)

func Create(c *gin.Context) {
	room := models.PrivateRoom{}
	err := c.ShouldBindBodyWith(&room, binding.JSON)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	theRoom := new(models.PrivateRoom)
	id := room.ID
	if id != 0 {
		database.MysqlDataBase.First(&theRoom, id)
		c.JSON(http.StatusOK, theRoom)
	} else {
		database.MysqlDataBase.Create(room)
		c.JSON(http.StatusOK, room)
	}
}

package friend

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"metocs_go/models"
	"metocs_go/utils"
	"strconv"
)

// AddFriend 用户添加一个好友
func AddFriend(c *gin.Context) {
	friend := &models.Friend{}
	err := c.ShouldBindBodyWith(friend, binding.JSON)
	if err != nil {
		log.Println(err.Error())
		return
	}
	claim := utils.GetClaim(c)
	friend.UserId = claim.ID
	friend.AddFriend()
	utils.SuccessNoMessage(c)
}

//GetFriends 获取用户好友列表
func GetFriends(c *gin.Context) {
	c.DefaultQuery("pageSize", "10")
	c.DefaultQuery("pageNum", "1")

	//解析int 参数
	value := c.Query("pageNum")
	pageNum, _ := strconv.Atoi(value)
	valueSize := c.Query("pageSize")
	pageSize, _ := strconv.Atoi(valueSize)

	claim := utils.GetClaim(c)
	var friends []models.Friend
	utils.Page(pageSize, pageNum).Where("user_id = ?", claim.ID).Find(&friends).Order("mark_name")
	utils.Success(c, friends)
}

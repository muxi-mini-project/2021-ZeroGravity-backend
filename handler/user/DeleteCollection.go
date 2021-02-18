package user

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

//DeleteCollection  is used to cancel user favorites 取消用户收藏
func DeleteCollection(c *gin.Context) {
	var decreaseCollection model.IdeaModel
	if err := c.BindJSON(&decreaseCollection); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
}

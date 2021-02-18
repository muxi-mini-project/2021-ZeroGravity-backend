package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

//DeleteComment  is used to delete comments  删除评论
func DeleteComment(c *gin.Context) {
	var decreaseComment model.IdeaModel
	if err := c.BindJSON(&decreaseComment); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
}

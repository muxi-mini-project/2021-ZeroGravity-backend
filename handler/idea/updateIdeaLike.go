package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

// UpdateIdeaLike is used to add a like record of idea 新增想法点赞记录
func UpdateIdeaLike(c *gin.Context) {
	var increaseIdeaLike model.IdeaLikeModel
	if err := c.BindJSON(&increaseIdeaLike); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
}

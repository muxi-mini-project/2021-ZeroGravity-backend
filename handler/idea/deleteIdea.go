package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

// DeleteIdea is used to delete ideas 删除想法
func DeleteIdea(c *gin.Context) {
	var decreaseIdea model.IdeaModel
	if err := c.BindJSON(&decreaseIdea); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
}

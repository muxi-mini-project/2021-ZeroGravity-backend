package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

// CreateIdea is used to post ideas 新增想法
func CreateIdea(c *gin.Context) {
	var increaseIdea model.IdeaModel
	if err := c.BindJSON(&increaseIdea); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
}

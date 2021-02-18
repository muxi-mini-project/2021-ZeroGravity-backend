package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

// CreateComment is used to post comments 新增评论
func CreateComment(c *gin.Context) {
	var increasecomment model.CommentModel
	if err := c.BindJSON(&increasecomment); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}

	// 调用服务
}

package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

// UpdateCommentLike is used to add a like record of comment 新增评论点赞记录
func UpdateCommentLike(c *gin.Context) {
	var increaseCommentLike model.CommentLikeModel
	if err := c.BindJSON(&increaseCommentLike); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
}

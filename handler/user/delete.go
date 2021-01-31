package user

import (
	"github.com/2021-ZeroGravity-backend/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)





//DecreaseIdea is used to delete ideas 删除想法

func DecreaseIdea (c *gin.Context) {
	var  decreaseIdea model.Idea
	token := c.Request.Header.Get("token")
	claims, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}
	if err := c.BindJSON(&decreaseIdea); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
 DeleteResult:= model.DeleteIdea (decreaseIdea)
 c.JSON(200, gin.H{"DeleteResult": DeleteResult})

}












//DecreaseComment  is used to delete comments  删除评论














//DecreaseIdeaLike  is used to delete like records of ideas  删除想法点赞记录













//DecreaseCommentLike  is used to delete like records of comments 删除评论点赞记录















//DecreaseCollection  is used to cancel user favorites 取消用户收藏



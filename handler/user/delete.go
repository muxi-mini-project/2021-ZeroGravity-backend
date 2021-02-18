package user

import (
	"github.com/2021-ZeroGravity-backend/model"

	"github.com/gin-gonic/gin"
)

//DecreaseIdea is used to delete ideas 删除想法
func DecreaseIdea(c *gin.Context) {
	var decreaseIdea model.Idea
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&decreaseIdea); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	DeleteResult := model.DeleteIdea(decreaseIdea)
	c.JSON(200, gin.H{"DeleteResult": DeleteResult})

}

//DecreaseComment  is used to delete comments  删除评论
func DecreaseComment(c *gin.Context) {
	var decreaseComment model.Idea
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&decreaseComment); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	DeleteResult := model.DeleteIdea(decreaseComment)
	c.JSON(200, gin.H{"DeleteResult": DeleteResult})

}

//DecreaseIdeaLike  is used to delete like records of ideas  删除想法点赞记录
func DecreaseIdeaLike(c *gin.Context) {
	var decreaseIdeaLike model.Idea
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&decreaseIdeaLike); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	DeleteResult := model.DeleteIdea(decreaseIdeaLike)
	c.JSON(200, gin.H{"DeleteResult": DeleteResult})

}

//DecreaseCommentLike  is used to delete like records of comments 删除评论点赞记录
func DecreaseCommentLike(c *gin.Context) {
	var decreaseCommentLike model.Idea
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&decreaseCommentLike); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	DeleteResult := model.DeleteIdea(decreaseCommentLike)
	c.JSON(200, gin.H{"DeleteResult": DeleteResult})

}

//DecreaseCollection  is used to cancel user favorites 取消用户收藏
func DecreaseCollection(c *gin.Context) {
	var decreaseCollection model.Idea
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&decreaseCollection); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	DeleteResult := model.DeleteIdea(decreaseCollection)
	c.JSON(200, gin.H{"DeleteResult": DeleteResult})

}

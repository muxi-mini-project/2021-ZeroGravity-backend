package user

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

// Register creates a new user account.新增用户
func Register(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	Nickname := model.NewUser(user)
	c.JSON(200, gin.H{"Nickname": Nickname})

}

//IncreaseComment is used to post comments 新增评论
func IncreaseComment(c *gin.Context) {
	var increasecomment model.Comment
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&increasecomment); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	Content := model.NewComment(increasecomment)
	c.JSON(200, gin.H{"Content": Content})

}

//IncreaseIdea is used to post ideas 新增想法
func IncreaseIdea(c *gin.Context) {
	var increaseIdea model.Idea
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&increaseIdea); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	Content := model.NewIdea(increaseIdea)
	c.JSON(200, gin.H{"Content": Content})

}

//IncreaseIdeaLike is used to add a like record of idea 新增想法点赞记录
func IncreaseIdeaLike(c *gin.Context) {
	var increaseIdeaLike model.IdeaLikeRecord
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&increaseIdeaLike); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	BelikedId := model.NewIdeaLike(increaseIdeaLike)
	c.JSON(200, gin.H{"BelikedId": BelikedId})

}

//IncreaseCommentLike is used to add a like record of comment 新增评论点赞记录
func IncreaseCommentLike(c *gin.Context) {
	var increaseCommentLike model.CommentLikeRecord
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&increaseCommentLike); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	BelikedId := model.NewCommentLike(increaseCommentLike)
	c.JSON(200, gin.H{"BelikedId": BelikedId})

}

//IncreaseCollection  is used to add a collection record of idea 新增收藏记录
func IncreaseCollection(c *gin.Context) {
	var increaseCollection model.Collection
	/*token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}*/
	if err := c.BindJSON(&increaseCollection); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	id := model.NewCollection(increaseCollection)
	c.JSON(200, gin.H{"id": id})

}

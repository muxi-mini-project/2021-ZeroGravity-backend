package user

import (
	"github.com/2021-ZeroGravity-backend/model"
	"log"
	"errors"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)


// Register creates a new user account.新增用户
func Register (c *gin.Context) {
	var  user model.User
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	Nickname:= model.NewUser(user)
    c.JSON(200, gin.H{"Nickname": Nickname})

}


	



//Login is used for user login 用户登录
    func Login (c *gin.Context) {
	    var  login  model.User

}











//IncreaseComment is used to post comments 新增评论
    func IncreaseComment (c *gin.Context) {
		var  increasecomment  model.Comment
		token := c.Request.Header.Get("token")
		claims, err := model.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"message": err.Error()})
			return
		}
		if err := c.BindJSON(&increasecomment); err != nil {
			c.JSON(400, gin.H{"message": "输入有误，格式错误"})
			return
		}
	id:= model.NewComment(increasecomment)
	c.JSON(200, gin.H{"id": id})

}













//IncreaseIdea is used to post ideas 新增想法
    func IncreaseIdea (c *gin.Context) {
		var  increaseIdea model.Idea
		token := c.Request.Header.Get("token")
		claims, err := model.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"message": err.Error()})
			return
		}
		if err := c.BindJSON(&increaseidea); err != nil {
			c.JSON(400, gin.H{"message": "输入有误，格式错误"})
			return
		}
	idea_id:= model.NewIdeaLike(increaseIdea)
	c.JSON(200, gin.H{"idea_id": idea_id})

}
  




//IncreaseIdeaLike is used to add a like record of idea 新增想法点赞记录
    func IncreaseIdeaLike (c *gin.Context) {
	    var  increaseIdeaLike  model.IdeaLikeRecord
		token := c.Request.Header.Get("token")
		claims, err := model.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"message": err.Error()})
			return
		}
		if err := c.BindJSON(&increaseIdeaLike); err != nil {
			c.JSON(400, gin.H{"message": "输入有误，格式错误"})
			return
		}
	id:= model.NewIdeaLike(ideaLikeRecord)
	c.JSON(200, gin.H{"id": id})

}




//IncreaseCommentLike is used to add a like record of comment 新增评论点赞记录
    func IncreaseCommentLike (c *gin.Context) {
	    var  increaseCommentLike  model.CommentLikeRecord
		token := c.Request.Header.Get("token")
		claims, err := model.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"message": err.Error()})
			return
		}
		if err := c.BindJSON(&increaseCommentLike); err != nil {
			c.JSON(400, gin.H{"message": "输入有误，格式错误"})
			return
		}
	id:= model.NewCommentLike(increaseCommentLike)
	c.JSON(200, gin.H{"id": id})

}




//IncreaseCollection  is used to add a collection record of idea 新增收藏记录
    func IncreaseCollection(c *gin.Context) {
	var  increaseCollection  model.Collection
    token := c.Request.Header.Get("token")
		claims, err := model.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"message": err.Error()})
			return
		}
		if err := c.BindJSON(&increaseCollection); err != nil {
			c.JSON(400, gin.H{"message": "输入有误，格式错误"})
			return
		}
	id:= model.NewCollection(increaseCollection)
	c.JSON(200, gin.H{"id": id})

}


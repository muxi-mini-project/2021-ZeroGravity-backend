package user

import (
	"github.com/gin-gonic/gin"
)

// UserCollection  is used to return to the front-end user's own favorite ideas 用来返回给前端用户自己收藏的想法
func UserCollection(c *gin.Context) {
	// var userCollection model.Collection

}

// UserIdea is used to return to front-end users to publish their own ideas 用来返回给前端用户自己发布的想法
func UserIdea(c *gin.Context) {
	// var userIdea model.Idea

}

// UserInfo is used return personal information to front-end users  用来返回给前端用户的个人信息
func UserInfo(c *gin.Context) {
	// var userInfo model.User

}

// UserComment  is used to return user comments to the front end  用来返回用户发布的评论给前端
func UserComment(c *gin.Context) {
	// var userComment model.Comment

}

//IdeaSearch  is used to return the queried idea information to the front end 返回查询到的想法信息给前端
func IdeaSearch(c *gin.Context) {
	// var ideaSearch model.Idea

}

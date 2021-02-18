package router

import (
	"github.com/2021-ZeroGravity-backend/handler/sd"
	"github.com/2021-ZeroGravity-backend/handler/user"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// middleware

	// router
	g.POST("/api/v1/register", user.Register)
	g.POST("/api/v1/login", user.Login)

	g1 := g.Group("/api/v1/user")
	{
		g1.POST("/:user_id/increase_collection", user.IncreaseCollection)
		g1.GET("/:user_id/collection", user.UserCollection)
		g1.GET("/:user_id/idea", user.UserIdea)
		g1.GET("/:user_id/", user.UserInfo)
		g1.GET("/:user_id/comment", user.UserComment)
		// g1.PUT("/:user_id/information", user.ChangeUserInfo)
		g1.DELETE("/:id/cancel_collection", user.DecreaseCollection)
	}
	g2 := g.Group("/api/v1/idea")
	{
		g2.POST("/:user_id/increase", user.IncreaseIdea)
		g2.DELETE("/:idea_id/reduction", user.DecreaseIdea)

	}
	g3 := g.Group("/api/v1/comment")
	{
		g3.POST("/:user_id/increase", user.IncreaseComment)
		g3.DELETE("/:id/reduction", user.DecreaseComment)
	}
	g4 := g.Group("/api/v1/like")
	{
		g4.DELETE("/:id/idea_record", user.DecreaseIdeaLike)
		g4.DELETE("/:id/comment_record", user.DecreaseCommentLike)
		g4.POST("/:user_id/idea", user.IncreaseIdeaLike)
		g4.POST("/:user_id/comment", user.IncreaseCommentLike)

	}
	g5 := g.Group("/api/v1/search")
	{
		g5.GET("/:user_id/idea_search", user.IdeaSearch)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

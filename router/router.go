package router

import (
	"net/http"

	"github.com/2021-ZeroGravity-backend/handler/idea"
	"github.com/2021-ZeroGravity-backend/handler/sd"
	"github.com/2021-ZeroGravity-backend/handler/user"
	"github.com/2021-ZeroGravity-backend/router/middleware"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// middleware
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// router
	g.POST("/api/v1/register", user.Register)
	g.POST("/api/v1/login", user.Login)

	g1 := g.Group("/api/v1/user")
	g1.Use(middleware.AuthMiddleware)
	{
		g1.GET("/detail/:id", user.GetUserInfo)
		g1.PUT("", user.UpdateUserInfo)
		// g1.GET("/notice",user.GetNotice)
	}

	g2 := g.Group("/api/v1/idea")
	g2.Use(middleware.AuthMiddleware)
	{
		g2.POST("", idea.CreateIdea)
		g2.DELETE("detail/:idea_id", idea.DeleteIdea)
		g2.GET("/collection", idea.GetCollection)
		g2.GET("/like", idea.GetIdeaLike)
		g2.POST("/collection", idea.CreateCollection)
		g2.DELETE("/collection", idea.DeleteCollection)
		g2.GET("/list", idea.GetIdeaList)
		g2.GET("/detail/:id/comment", idea.GetCommentList)
		// g2.GET("/detail/:id",idea.GetIdea) 获取单个 idea
	}

	g3 := g.Group("/api/v1/comment")
	g3.Use(middleware.AuthMiddleware)
	{
		g3.POST("", idea.CreateComment)
		g3.DELETE("/:comment_id", idea.DeleteComment)
	}

	g4 := g.Group("/api/v1/like")
	g4.Use(middleware.AuthMiddleware)
	{
		g4.PUT("/idea/:idea_id", idea.UpdateIdeaLike)
		g4.PUT("/comment/:comment_id", idea.UpdateCommentLike)

	}

	/*g5 := g.Group("/api/v1/search")
	g5.Use(middleware.AuthMiddleware)
	{
		g5.GET("", user.Search)
	}*/

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

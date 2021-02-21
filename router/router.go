package router

import (
	"net/http"

	"github.com/2021-ZeroGravity-backend/handler/auth"
	"github.com/2021-ZeroGravity-backend/handler/collection"
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
	authRouter := g.Group("/api/v1/auth")
	{
		authRouter.POST("/api/v1/register", auth.Register)
		authRouter.POST("/api/v1/login", auth.Login)
	}

	userRouter := g.Group("/api/v1/user")
	userRouter.Use(middleware.AuthMiddleware)
	{
		userRouter.GET("/detail/:id", user.GetUserInfo)
		userRouter.PUT("", user.UpdateUserInfo)
		// userRouter.GET("/notice",user.GetNotice)
	}

	ideaRouter := g.Group("/api/v1/idea")
	ideaRouter.Use(middleware.AuthMiddleware)
	{
		// idea
		ideaRouter.POST("", idea.CreateIdea)
		ideaRouter.DELETE("detail/:idea_id", idea.DeleteIdea)
		ideaRouter.GET("/list", idea.GetIdeaList)
		// ideaRouter.GET("/detail/:id",idea.GetIdea) 获取单个 idea

		ideaRouter.GET("/liked", idea.GetIdeaLike)

		// comment
		ideaRouter.POST("/comment", idea.CreateComment)
		ideaRouter.DELETE("/comment/:comment_id", idea.DeleteComment)
		ideaRouter.GET("/detail/:id/comment", idea.GetCommentList)

		// like
		ideaRouter.PUT("/detail/:idea_id/like", idea.UpdateIdeaLike)
		ideaRouter.PUT("/comment/:comment_id/like", idea.UpdateCommentLike)
	}

	collectionRouter := g.Group("/api/v1/collection")
	collectionRouter.Use(middleware.AuthMiddleware)
	{
		collectionRouter.GET("/collection", collection.GetCollection)
		collectionRouter.POST("/collection", collection.CreateCollection)
		collectionRouter.DELETE("/collection", collection.DeleteCollection)
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

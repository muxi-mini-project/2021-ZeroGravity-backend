package router
  
  import (
	  
	"github.com/gin-gonic/gin"
  )
 func Router (r *gin.Engine)  {
	 
 
	router := gin.Default()


	router.POST("/api/v1/register", Handler.Register)
	router.POST("/api/v1/login", Handler.Login)

	g1 := router.Group("/api/v1/user")
	{
		g1.POST("/:user_id/increase_collection", Handler.IncreaseCollection)
		g1.GET("/:user_id/collection",Handler.UserCollection)
		g1.GET("/:user_id/idea", Handler.UserIdea)
		g1.GET("/:user_id/", Handler.UserInfo)
		g1.GET("/:user_id/comment", Handler.UserComment)
		g1.PUT("/:user_id/information", Handler.ChangeUserInfo)
		g1.DELETE("/:user_id/cancel_collection",Handler.DeleteCollection)
	}	
    g2 := router.Group("/api/v1/idea")
	{
		g2.POST("/:user_id/increase", Handler.IncreaseIdea)
		g2.DELETE("/:user_id/reduction", Handler.DeleteIdea)
		
	}	
	g3 := router.Group("/api/v1/comment"){
        g3.POST("/:user_id/increase", Handler.IncreaseComment)
		g3.DELETE("/:user_id/reduction", Handler.DeleteComment)
    }
	g4 := router.Group("/api/v1/like"){
        g4.DELETE("/:user_id/idea_record", Handler.DeleteIdeaLike)
		g4.DELETE("/:user_id/comment_record", Handler.DeleteCommentLike)
		g4.POST("/:user_id/idea", Handler.IncreaseIdeaLike)
		g4.POST("/:user_id/comment", Handler.IncreaseCommentLike)

	}
	g5 :=router.Group("/api/v1/search"){
		g5.GET("/:user_id/idea_search", Handler.IdeaSearch) 
    }
	

    router.Run(":9091")
}
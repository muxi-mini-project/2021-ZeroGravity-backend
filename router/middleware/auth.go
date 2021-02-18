package middleware

import (
	"github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/pkg/auth"
	"github.com/2021-ZeroGravity-backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ... 认证中间件
func AuthMiddleware(c *gin.Context) {
	// Parse the json web token.
	ctx, err := auth.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, errno.ErrTokenInvalid, err.Error())
		c.Abort()
		return
	}

	c.Set("userID", ctx.ID)
	c.Set("expiresAt", ctx.ExpiresAt)

	c.Next()
}

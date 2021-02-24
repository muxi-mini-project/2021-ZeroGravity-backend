package auth

import (
	"encoding/base64"

	"github.com/2021-ZeroGravity-backend/handler"
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/user"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login ... 登陆
func Login(c *gin.Context) {
	log.Info("login function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req LoginRequest
	// 获取账号和密码
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	decodePassword, err := base64.StdEncoding.DecodeString(req.AccountPassword)
	if err != nil {
		handler.SendError(c, errno.ErrDecoding, nil, err.Error(), "")
		return
	}

	// 调用服务
	var token string
	token, err = user.Login(req.Account, string(decodePassword))
	if err != nil {
		SendError(c, errno.ErrPasswordIncorrect, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, &LoginResponse{
		Token: token,
	})
}

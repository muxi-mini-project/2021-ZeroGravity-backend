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

// @Summary 用户输入账号和密码登录 
// @Description 登录
// @Tags auth
// @Accept  json
// @Produce  json
// @Param req body auth.LoginRequest true " Account 账户   AccountPassword 密码   "
// @Success 200 {object} LoginResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/auth/api/v1/login [post]
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
	token, userId, err := user.Login(req.Account, string(decodePassword))
	if err != nil {
		SendError(c, errno.ErrPasswordIncorrect, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, &LoginResponse{
		Token:  token,
		UserId: userId,
	})
}

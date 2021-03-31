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
// @Summary  用户输入账号和密码登录
// @Description 用户注册
// @Tags auth
// @Accept  json
// @Produce  json
// @Param req body auth.CreateUserRequest true "Account--账户 AccountPassword--账户密码	NickName--昵称	Avatar--头像 "
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/auth/api/v1/register [post]
// Register creates a new user account ... 新增用户
func Register(c *gin.Context) {
	log.Info("User register function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// base64解密
	decodePassword, err := base64.StdEncoding.DecodeString(req.AccountPassword)
	if err != nil {
		handler.SendError(c, errno.ErrDecoding, nil, err.Error(), "")
		return
	}

	// 调用服务
	if err := user.Register(req.Account, string(decodePassword), req.NickName, req.Avatar); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

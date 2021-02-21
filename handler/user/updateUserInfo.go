package user

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/user"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

// UpdateUserInfo ... 更改用户的信息
func UpdateUserInfo(c *gin.Context) {
	log.Info("user getUserInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 从 token 获取
	uid := c.MustGet("userID").(int)

	// 获取请求
	var req UpdateUserInfoRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	if err := user.UpdateUserInfo(uid, req.Avatar, req.NickName); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

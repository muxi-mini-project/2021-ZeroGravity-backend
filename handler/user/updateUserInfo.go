package user

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/user"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap")

// @Summary 更改用户信息
// @Description 更改用户信息
// @Tags user
// @Accept  json
// @Produce  json
// @Param  token header string true  "UserId--用户ID"
// @Param  req body user.UpdateUserInfoRequest true  "Avatar头像|| NickName昵称"
// @Success  200 "成功"
// @Router /api/v1/user [put]
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

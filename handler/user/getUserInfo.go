package user

import (
	"strconv"

	"github.com/2021-ZeroGravity-backend/model"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)
// @Summary 得到用户信息  
// @Description 得到用户所有的个人信息
// @Tags user
// @Accept  json
// @Produce  json
// @Param  id path int true "userId--用户ID"
// @Success 200 {object}  user.GetUserInfoResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /api/v1/user/detail/:id [get]
// GetUserInfo ... 获取某个用户的信息
func GetUserInfo(c *gin.Context) {
	log.Info("user getUserInfo function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var uid int
	var err error

	// 从路由获取参数
	uid, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	item, err := model.GetUserById(uid)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &GetUserInfoResponse{
		Id:       item.Id,
		Avatar:   item.Avatar,
		NickName: item.NickName,
		Energy:   item.Energy,
	}

	SendResponse(c, errno.OK, resp)
}

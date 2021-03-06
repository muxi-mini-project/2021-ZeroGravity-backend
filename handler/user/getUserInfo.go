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
// @Summary Get information about a certain user  
// @Description All information returned to a user on the front end
// @Tags user
// @Accept  json
// @Produce  json
// @Param  id body string true  "All information returned to a user on the front end"
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

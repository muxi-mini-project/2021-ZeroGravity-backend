package message

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 获取消息提示
// @Description 获取消息提示
// @Tags message
// @Accept  json
// @Produce  json
// @Param token header string true  "后端给前端的Token"
// @Success 200 {object}  message.GetMessageTipResponse  "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/message/tip [get]
// GetMessageTip ... 获取消息提示
func GetMessageTip(c *gin.Context) {
	log.Info("Message getMessageTip function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	uid := c.MustGet("userID").(int)

	// service
	count, err := model.GetMessageTipByUserId(uid)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &GetMessageTipResponse{}
	if count != 0 {
		resp.HaveMessage = true
	}

	SendResponse(c, errno.OK, resp)
}

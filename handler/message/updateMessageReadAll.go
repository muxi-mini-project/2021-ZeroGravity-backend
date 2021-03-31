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
// @Summary   修改信息为已读  
// @Description  修改信息为已读
// @Accept  json
// @Produce  json
// @Param token header string true  "后端给前端的token"
// @Success 200 "成功" 
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/message/readall [put]
// UpdateMessageReadAll ... 修改信息为已读
func UpdateMessageReadAll(c *gin.Context) {
	log.Info("Message getMessageTip function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	uid := c.MustGet("userID").(int)

	// service
	if err := model.ReadAll(uid); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

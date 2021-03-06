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
// @Summary Modify the information as read  
// @Description Change the status of the messages that the user has viewed to read
// @Tags message
// @Accept  json
// @Produce  json
// @Param userID head string true  "Change the status of the messages that the user has viewed to read"
// @Success 200 "成功" 
// @Router/api/v1/message/readall [put]

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

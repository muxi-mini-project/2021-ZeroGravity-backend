package message

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
// @Summary Get user comment information  
// @Description Information returned to front-end user reviews
// @Tags message
// @Accept  json
// @Produce  json
// @Param token header string true  "Information returned to front-end user information"
// @Param limit query  int true  "limit"
// @Param page  query  int true  "page"
// @Success 200 {object}  message.GetMessageForCommentResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router/api/v1/message/comment [get]
// GetMessageForComment ... 获取评论信息
func GetMessageForComment(c *gin.Context) {
	log.Info("Message getMessageForComment function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	uid := c.MustGet("userID").(int)

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	item, err := model.GetMessageForComment(uid, page*limit, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	var resp GetMessageForCommentResponse
	resp.Count = len(item)
	resp.List = item

	SendResponse(c, errno.OK, resp)
}

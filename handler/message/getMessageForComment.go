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
// @Summary   拉取评论我的消息
// @Description  拉取评论我的消息
// @Tags message
// @Accept  json
// @Produce  json
// @Param token header string true  "后端给前端的Token"
// @Param limit query  int true  "limit--偏移量指定开始返回记录之前要跳过的记录数"
// @Param page  query  int true  "page--限制指定要检索的记录数"
// @Success 200 {object}  message.GetMessageForCommentResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/message/comment [get]
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

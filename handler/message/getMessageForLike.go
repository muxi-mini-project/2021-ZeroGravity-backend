package message

import (
	"strconv"

	"github.com/2021-ZeroGravity-backend/model"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap")
// @Summary 获取点赞我的信息
// @Description  获取点赞我的信息
// @Tags message
// @Accept  json
// @Produce  json
// @Param token header string true  "Like information returned to front-end users"
// @Param limit query  int true  "limit--偏移量指定开始返回记录之前要跳过的记录数"
// @Param page  query  int true  "page--限制指定要检索的记录数"
// @Success 200 {object}  message.GetMessageForLikeResponse  "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/message/like [get])
// GetMessageForLike ... 获取点赞信息
func GetMessageForLike(c *gin.Context) {
	log.Info("Message getMessageForLike function called.",
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

	item, err := model.GetMessageForLike(uid, page*limit, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	var resp GetMessageForLikeResponse
	resp.Count = len(item)
	resp.List = item

	SendResponse(c, errno.OK, resp)
}

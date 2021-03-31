package collection

import (
	"strconv"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/2021-ZeroGravity-backend/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary  查询用户的收藏列表
// @Description 查看用户的收藏列表
// @Tags collection
// @Accept  json
// @Produce  json
// @Param token header string true  "userid用户ID"
// @Param limit query  int true "limit--偏移量指定开始返回记录之前要跳过的记录数 "
// @Param page  query  int true "page--限制指定要检索的记录数 "
// @Success 200 {object} collection.IdeaResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/collection [get]
// GetCollection ... 获取某个用户的收藏列表
func GetCollection(c *gin.Context) {
	log.Info("Idea getCollection function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var userid, limit, page int
	var err error

	// 从 token 获取 userid
	userid = c.MustGet("userID").(int)

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	item, err := idea.GetCollection(userid, page*limit, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &IdeaResponse{}
	resp.Count = len(item)
	resp.List = item

	SendResponse(c, errno.OK, resp)
}

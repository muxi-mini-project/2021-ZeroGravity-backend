package idea

import (
	"strconv"

	"github.com/2021-ZeroGravity-backend/service/idea"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 发布的想法下对应的评论  
// @Description 想法的评论列表
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "uid--用户ID"
// @Param id path int true "id--被评论的想法ID"
// @Param limit query  int true "limit--偏移量指定开始返回记录之前要跳过的记录数"
// @Param page  query  int true "page--限制指定要检索的记录数 "
// @Success 200 {object}  idea.CommentResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/detail/:id/comment [get]
// GetCommentList ... 获取某个想法的评论列表
func GetCommentList(c *gin.Context) {
	log.Info("idea getComment function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var id, uid, limit, page int
	var err error

	// 从 token 获取 userid
	uid = c.MustGet("userID").(int)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

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
	item, err := idea.GetCommentList(id, uid, page*limit, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &CommentResponse{}
	resp.Count = len(item)
	resp.List = item

	SendResponse(c, errno.OK, resp)
}

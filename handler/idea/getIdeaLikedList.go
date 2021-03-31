package idea

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

// @Summary 获取用户个人发布的想法被点赞的情况  
// @Description  获取想法被点赞的信息表
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "UserId--用户ID"
// @Param limit query  int true "limit--偏移量指定开始返回记录之前要跳过的记录数"
// @Param page  query  int true "page--限制指定要检索的记录数 "
// @Success 200 {object}  idea.IdeaResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/liked [get]
// GetIdeaLikedList ... 获取某个用户点赞的想法列表
func GetIdeaLikedList(c *gin.Context) {
	log.Info("Idea getIdeaLike function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var id, limit, page int
	var err error

	// 从 token 获取 userid
	id = c.MustGet("userID").(int)

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
	item, err := idea.GetIdeaLikedList(id, page*limit, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &IdeaResponse{}
	resp.Count = len(item)
	resp.List = item

	SendResponse(c, errno.OK, resp)
}

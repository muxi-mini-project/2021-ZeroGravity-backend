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

// @Summary  进入空间后看到的想法
// @Description  显示用户发布的想法
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "uid--用户ID"
// @Param limit query  int true "limit--偏移量指定开始返回记录之前要跳过的记录数"
// @Param page  query  int true "page--限制指定要检索的记录数 "
// @Param userId query int true "userId--用户ID"
// @Param privicy query int true "privicy 获取想法的私有策略 0->不获取私有 1->获取私有 默认不获取 || 若获取则默认为获取自己的想法，需要服务判断 uid"
// @Param index query  int true "index--获取排序规则 0->默认时间排序 1->热度排序"
// @Success 200 {object}  idea.IdeaResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /api/v1/idea/list [get]
// GetIdeaList ... 获取想法列表
func GetIdeaList(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var uid, limit, page, userId, index, privacy, space int
	var err error

	// 从 token 获取 userid
	uid = c.MustGet("userID").(int)

	// 获取空间选择，默认 0 获取全部
	space, err = strconv.Atoi(c.DefaultQuery("space", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
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

	// 获取指定的 userid 不指定为默认 0
	userId, err = strconv.Atoi(c.DefaultQuery("userId", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取想法的私有策略 0->不获取私有 1->获取私有 默认不获取
	// 若获取则默认为获取自己的想法，需要服务判断 uid
	privacy, err = strconv.Atoi(c.DefaultQuery("privacy", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 获取排序规则 0->默认时间排序 1->热度排序
	index, err = strconv.Atoi(c.DefaultQuery("index", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	item, err := idea.GetIdeaList(uid, page*limit, limit, privacy, index, userId, space)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &IdeaResponse{}
	resp.Count = len(item)
	resp.List = item

	SendResponse(c, errno.OK, resp)
}

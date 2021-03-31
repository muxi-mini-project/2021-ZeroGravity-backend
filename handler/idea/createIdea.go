package idea

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 发布想法
// @Description 新增想法
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "userid--用户ID"
// @Param req body idea.CreateIdeaRequest true  "Content 想法内容 || Space  1->情绪 2->脑洞 3->沉思 4->探梦 || Privacy  0->公开 1->隐私"
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea [post]
// CreateIdea is used to post ideas 新增想法
func CreateIdea(c *gin.Context) {
	log.Info("Create Idea function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	var userid int

	// 从 token 获取 userid
	userid = c.MustGet("userID").(int)
	var req CreateIdeaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// space 和 privacy 判断
	if req.Space == 0 && req.Privacy == 0 {
		SendBadRequest(c, errno.ErrSpace, nil, "idea must choose a space", GetLine())
	}

	// 调用服务
	if err := idea.CreateIdea(req.Content, userid, req.Space, req.Privacy); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)

}

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

// @Summary Add ideas
// @Description Add a thought record to the database
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "userid"
// @Param req body idea.CreateIdeaRequest true  "Add a thought record to the database"
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

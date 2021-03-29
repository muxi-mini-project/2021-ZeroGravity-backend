package idea

import (
	"strconv"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
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
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/detail/:id [delete]
// DeleteIdea is used to delete ideas 删除想法
func DeleteIdea(c *gin.Context) {

	log.Info("Delete Idea function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	var userid int

	// 从 token 获取 userid
	userid = c.MustGet("userID").(int)

	IdeaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}
	Idea, err := model.GetIdeaById(IdeaId)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	if Idea.UserId != userid {
		SendError(c, errno.ErrMatch, nil, "", GetLine())
		return
	}
	// 调用服务
	if err := model.DeleteIdea(IdeaId, userid); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)

}

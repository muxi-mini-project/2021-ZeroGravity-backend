package report

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/report"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
// @Summary 举报用户 
// @Description 举报用户
// @Tags report
// @Accept  json
// @Produce  json
// @Param  req body report.CreateReportRequest true  "UserID--被举报的用户ID || Kind-- ||   Reason--原因   || CommentId --评论的ID|| ideaId--想法的ID" 
// @Param token header string true  "userId--用户ID"
// @Success 200 "成功" 
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno  
// @Router /api/v1/report [post]
// CreateReport is used to post report
func CreateReport(c *gin.Context) {
	log.Info("Create Report function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateReportRequest
	if err := c.Bind(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	uid := c.MustGet("userID").(int)

	// 调用服务
	if err := report.CreatReport(uid, req.UserId, req.Kind, req.CommentId, req.IdeaId, req.Reason); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

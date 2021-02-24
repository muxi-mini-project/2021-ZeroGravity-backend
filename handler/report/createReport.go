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

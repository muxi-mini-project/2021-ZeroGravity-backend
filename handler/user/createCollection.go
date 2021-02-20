package user

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/user"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCollection  is used to add a collection record of idea 新增收藏记录
func CreateCollection(c *gin.Context) {
	log.Info("Create Collection function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req CreateCollectionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return

	}
	// 调用服务
	if err := user.CreateCollection(req.IdeaId, req.CollectorId); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

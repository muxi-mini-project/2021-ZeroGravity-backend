package collection

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
// @Summary Add favorite record  
// @Description Add a new favorite record to the database
// @Tags collection
// @Accept  json
// @Produce  json
// @Param req body collection.CreateCollectionRequest  true "Add a new favorite record to the database "
// @Success 200 "成功"
// @Router /api/v1/collection [get]

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
	if err := idea.CreateCollection(req.IdeaId, req.CollectorId); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

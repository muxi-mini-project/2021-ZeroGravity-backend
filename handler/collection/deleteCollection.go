package collection

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
// @Summary Delete favorites   
// @Description Delete favorite records from the database 
// @Tags collection
// @Accept  json
// @Produce  json
// @Body  req body collection.DeleteCollectionRequest  true "Delete favorite records from the database  "
// @Success 200 "成功"
// @Router /api/v1/collection [delete]
// DeleteCollection  is used to add a collection record of idea 删除收藏记录
func DeleteCollection(c *gin.Context) {
	log.Info("Delete Collection function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req DeleteCollectionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return

	}
	// 调用服务
	if err := model.DeleteCollection(req.IdeaId, req.CollectorId); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)
}

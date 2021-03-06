package collection

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
// @Summary 添加收藏记录  
// @Description 新增收藏记录
// @Tags collection
// @Accept  json
// @Produce  json
// @Param token header string true  "userid--用户的ID"
// @Param req body collection.CreateCollectionRequest  true "IdeaId用户收藏的想法ID"
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/collection [post]
// CreateCollection  is used to add a collection record of idea 新增收藏记录
func CreateCollection(c *gin.Context) {
	log.Info("Create Collection function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var userid int

	// 从 token 获取 userid
	userid = c.MustGet("userID").(int)
	var req CreateCollectionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return

	}
	// 调用服务
	_, err := model.GetIdeaById(req.IdeaId)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	if err := idea.CreateCollection(req.IdeaId, userid); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, errno.OK, nil)
}

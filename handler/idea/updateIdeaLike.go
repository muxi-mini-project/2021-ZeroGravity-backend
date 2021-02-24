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

// UpdateIdeaLike is used to add a like record of idea 新增想法点赞记录
func UpdateIdeaLike(c *gin.Context) {

	log.Info("Update Idea Like function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req UpdateIdeaLikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	if err := idea.UpdateIdeaLike(req.IdeaId, req.LikersId, req.BelikedId); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	// TODO:likesum 增减和 creae message

	SendResponse(c, errno.OK, nil)

}

package idea

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/gin-gonic/gin"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/util"
	"go.uber.org/zap"

)

// UpdateCommentLike is used to add a like record of comment 新增评论点赞记录
func UpdateCommentLike(c *gin.Context) {
	
	log.Info("Update Comment Like function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	
	var req model.CreateCommentLikeRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	
	// 调用服务
	if err :=  idea.UpdateCommentLike(&req); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)

}

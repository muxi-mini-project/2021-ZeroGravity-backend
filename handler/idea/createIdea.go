package idea

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/2021-ZeroGravity-backend/service"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/util"
	"go.uber.org/zap"

)

// CreateIdea is used to post ideas 新增想法
func CreateIdea(c *gin.Context) {
	log.Info("Create Idea function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	
	var req model.CreateIdeaRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	
	// 调用服务
	if err := service.CreateIdea(&req); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, nil)

}
  

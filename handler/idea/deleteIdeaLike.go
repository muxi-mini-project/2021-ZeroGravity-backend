package idea

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	
	"github.com/gin-gonic/gin"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/util"
	"go.uber.org/zap"
)
//DeleteIdeaLike  is used to delete idea like 删除评论
func DeleteIdeaLike(c *gin.Context) {
	
	log.Info("Delete Comment like function called.",
	zap.String("X-Request-Id", util.GetReqID(c)))
		
	var req model.DeleteIdeaLikeRequest
		
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
		
	// 调用服务
	if err := idea.DeleteIdeaLike(&req); err != nil {
			SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
			return
		}
	
		SendResponse(c, errno.OK, nil)
	
	}


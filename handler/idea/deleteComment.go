package idea

import (
	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/model"

	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//DeleteComment  is used to delete comments  删除评论
func DeleteComment(c *gin.Context) {

	log.Info("Delete Comment function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req DeleteCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}
	uid := c.MustGet("userID").(int)
	// 调用服务
	if err := model.DeleteComment(req.Id, uid); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	// TODO:评论数减 1

	var i model.IdeaInfo
	model.DB.Self.Where("idea_id = ? ", req.Id).First(&i)
	i2 := i
	i2.CommentSum--
	model.DB.Self.Model(&i).Update(i2)
	SendResponse(c, errno.OK, nil)

}

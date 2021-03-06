package idea

import (
	"strconv"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/model"

	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 删除用户自己的评论
// @Description 删除评论
// @Tags idea
// @Accept  json
// @Produce  json
// @Param id path string true "id--评论的ID"
// @Param token header string true  "uid --用户的ID"
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/detail/:id/comment/:comment_id [delete]
//DeleteComment  is used to delete comments  删除评论
func DeleteComment(c *gin.Context) {

	log.Info("Delete Comment function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}
	uid := c.MustGet("userID").(int)
	CommentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	var com model.CommentModel
	err = model.DB.Self.Where("id = ? ", CommentId).First(&com).Error
	if err != nil {
		SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	if com.CommenterId != uid {
		SendError(c, errno.ErrMatch, nil, "", GetLine())
		return
	}

	// 调用服务
	if err := model.DeleteComment(CommentId, uid); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	var i model.IdeaModel
	model.DB.Self.Where("idea_id = ? ", id).First(&i)
	i.CommentSum--
	model.DB.Self.Model(&i).Where("idea_id = ? ", id).Update("comment_sum", i.CommentSum)
	SendResponse(c, errno.OK, nil)

}

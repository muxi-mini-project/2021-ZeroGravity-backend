package idea

import (
	"strconv"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/2021-ZeroGravity-backend/service/message"
	"github.com/2021-ZeroGravity-backend/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Add comment and like record
// @Description Add a comment and like record to the database
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "LikersId"
// @Param id path int true "CommentId"
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/comment/:id/like [put]
// UpdateCommentLike is used to add a like record of comment 新增评论点赞记录
func UpdateCommentLike(c *gin.Context) {

	log.Info("Update Comment Like function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	LikersId := c.MustGet("userID").(int)
	var req UpdateLikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	CommentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}
	if req.Choice > 2 || req.Choice < 1 {
		SendBadRequest(c, errno.ErrChoice, nil, "", GetLine())
		return
	}

	var COMment model.CommentModel
	result := model.DB.Self.Where("id = ? ", CommentId).First(&COMment)
	if result.Error != nil {
		SendError(c, errno.ErrDatabase, nil, result.Error.Error(), GetLine())
		return
	}
	//取消点赞
	if req.Choice == 2 {
		var comlike model.CommentLikeModel
		if result := model.DB.Self.Where("comment_id = ? AND likers_id = ? ", CommentId, LikersId).First(&comlike); result.Error == nil {
			//未点赞
			SendError(c, errno.ErrNotLike, nil, result.Error.Error(), GetLine())
			return
		} else {
			model.DB.Self.Delete(&comlike)
			var i model.CommentModel
			model.DB.Self.Where("id = ? ", CommentId).First(&i)
			i.LikesSum--
			model.DB.Self.Model(&i).Where("id = ? ", CommentId).Update("likes_sum", i.LikesSum)
			SendResponse(c, errno.OK, nil)
			return
		}
	}
	if req.Choice == 1 {
		if result := model.DB.Self.Where("comment_id = ? AND likers_id = ? ", CommentId, LikersId); result.Error != nil {
			//已点赞
			SendError(c, errno.ErrHaveLike, nil, result.Error.Error(), GetLine())
			return
		} else {
			// 调用服务
			if err := idea.UpdateCommentLike(CommentId, LikersId); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			var i model.CommentModel
			model.DB.Self.Where("id = ? ", CommentId).First(&i)
			i.LikesSum++
			model.DB.Self.Model(&i).Where("id = ? ", CommentId).Update("likes_sum", i.LikesSum)

			//creae message
			if err := message.CreateMessage(LikersId, COMment.CommenterId, 0, CommentId, 0, i.Content, ""); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			SendResponse(c, errno.OK, nil)
			return
		}
	}
}

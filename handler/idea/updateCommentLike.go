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

	var COMment model.CommentInfo
	result := model.DB.Self.Where("comment_id = ? ", CommentId).First(&COMment)
	if result.Error != nil {
		return
	}
	//取消点赞
	if req.Choice == 2 {
		var Idea model.IdeaLikeModel
		if result := model.DB.Self.Where("comment_id = ? AND likers_id = ? ", CommentId, LikersId).First(&Idea); result.Error != nil {
			//未点赞
			SendError(c, errno.ErrNotLike, nil, result.Error.Error(), GetLine())
			return
		} else {
			model.DB.Self.Delete(&Idea)
			var i model.CommentInfo
			model.DB.Self.Where("comment_id = ? ", CommentId).First(&i)
			i2 := i
			i2.LikesSum--
			model.DB.Self.Model(&i).Update(i2)
			SendResponse(c, errno.OK, nil)
			return
		}
	}
	if req.Choice == 1 {
		if result := model.DB.Self.Where("comment_id = ? AND likers_id = ? ", CommentId, LikersId); result.Error == nil {
			//已点赞
			SendError(c, errno.ErrHaveLike, nil, result.Error.Error(), GetLine())
			return
		} else {
			// 调用服务
			if err := idea.UpdateCommentLike(CommentId, LikersId, COMment.UserId); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			var i model.CommentInfo
			model.DB.Self.Where("comment_id = ? ", CommentId).First(&i)
			i2 := i
			i2.LikesSum++
			model.DB.Self.Model(&i).Update(i2)

			//creae message
			if err := message.CreateMessage(LikersId, COMment.UserId, 0, CommentId, 0, i.Content, ""); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			SendResponse(c, errno.OK, nil)
			return
		}
	}
}

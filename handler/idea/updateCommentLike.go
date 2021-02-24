package idea

import (
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

// UpdateCommentLike is used to add a like record of comment 新增评论点赞记录
func UpdateCommentLike(c *gin.Context) {

	log.Info("Update Comment Like function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var req UpdateCommentLikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	if req.Choice > 2 || req.Choice < 1 {
		c.JSON(400, gin.H{"message": "Fail.1 == 点赞， 2 == 取消点赞"})
		return
	}

	//取消点赞
	if req.Choice == 2 {
		var idea model.IdeaLikeModel
		if result := model.DB.Self.Where("comment_id = ? AND likers_id = ? ", req.CommentId, req.LikersId).First(&idea); result.Error != nil {
			//未点赞
			//SendError(c, errno.ErrDatabase, nil, nil, GetLine())
			c.JSON(400, gin.H{"message": "未点赞"})
			return
		} else {
			model.DB.Self.Delete(&idea)
			var i model.CommentInfo
			model.DB.Self.Where("comment_id = ? ", req.CommentId).First(&i)
			i2 := i
			i2.LikesSum--
			model.DB.Self.Model(&i).Update(i2)
			c.JSON(200, gin.H{"message": "取消成功"})
			return
		}
	}
	if req.Choice == 1 {

		if result := model.DB.Self.Where("comment_id = ? AND likers_id = ? ", req.CommentId, req.LikersId); result.Error == nil {
			//未点赞
			//SendError(c, errno.ErrDatabase, nil, nil, GetLine())
			c.JSON(400, gin.H{"message": "已点赞"})
			return
		} else {
			// 调用服务
			if err := idea.UpdateCommentLike(req.CommentId, req.LikersId, req.BelikedId); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			var i model.IdeaInfo
			//model.DB.Self.Where("idea_id = ? ", req.IdeaId).First(&i)
			i2 := i
			i2.LikesSum++
			model.DB.Self.Model(&i).Update(i2)

			//creae message
			if err := message.CreateMessage(req.LikersId, req.BelikedId, 0, req.CommentId, 0, i.Content, ""); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			c.JSON(200, gin.H{"message": "点赞成功"})
			return
		}
	}
	// 调用服务
	if err := idea.UpdateCommentLike(req.CommentId, req.LikersId, req.BelikedId); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	// TODO: likesum增减 和 create message

	SendResponse(c, errno.OK, nil)

}

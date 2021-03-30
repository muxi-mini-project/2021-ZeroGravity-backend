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

// @Summary Add thoughts like record
// @Description Add an idea and like record to the database
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "LikersId "
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/detail/:id/like [put]
// UpdateIdeaLike is used to add a like record of idea 新增想法点赞记录
func UpdateIdeaLike(c *gin.Context) {

	log.Info("Update Idea Like function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	LikersId := c.MustGet("userID").(int)
	var req UpdateLikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	IdeaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	if req.Choice > 2 || req.Choice < 1 {
		SendBadRequest(c, errno.ErrChoice, nil, "", GetLine())
		return
	}
	var IDEA model.IdeaModel
	result := model.DB.Self.Where("idea_id = ? ", IdeaId).First(&IDEA)
	if result.Error != nil {
		return
	}
	//取消点赞
	if req.Choice == 2 {
		var Idealike model.IdeaLikeModel
		if result := model.DB.Self.Where("idea_id = ? AND likers_id = ? ", IdeaId, LikersId).First(&Idealike); result.Error != nil {
			//未点赞
			SendError(c, errno.ErrNotLike, nil, result.Error.Error(), GetLine())
			return
		} else {
			model.DB.Self.Delete(&Idealike)
			var i model.IdeaModel
			model.DB.Self.Where("idea_id = ? ", IdeaId).First(&i)
			i.LikesSum--
			model.DB.Self.Model(&i).Where("idea_id = ? ", IdeaId).Update("likes_sum", i.LikesSum)
			SendResponse(c, errno.OK, nil)
			return
		}
	}
	if req.Choice == 1 {
		if result := model.DB.Self.Where("idea_id = ? AND likers_id = ? ", IdeaId, LikersId); result.Error != nil {
			//已点赞
			SendError(c, errno.ErrHaveLike, nil, result.Error.Error(), GetLine())
			return
		} else {
			// 调用服务
			if err := idea.UpdateIdeaLike(IdeaId, LikersId); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			var i model.IdeaModel
			model.DB.Self.Where("idea_id = ? ", IdeaId).First(&i)
			i.LikesSum++
			model.DB.Self.Model(&i).Where("idea_id = ? ", IdeaId).Update("likes_sum", i.LikesSum)

			//creae message
			if err := message.CreateMessage(LikersId, IDEA.PublisherId, 0, 0, IdeaId, i.Content, ""); err != nil {
				SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
				return
			}
			SendResponse(c, errno.OK, nil)
			return
		}
	}
}

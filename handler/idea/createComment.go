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

// @Summary Add comment
// @Description Add comment records to the database
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "userid--用户的ID"
// @Param req body idea.CreateCommentRequest true  "CommentedId 被评论者的ID || Content 评论内容"
// @Param idea_id path int true "IdeaId--想法ID"
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/detail/:id/comment [post]
// CreateComment is used to post comments 新增评论
func CreateComment(c *gin.Context) {

	log.Info("Create Comment function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	var userid int

	// 从 token 获取 userid
	userid = c.MustGet("userID").(int)
	var req CreateCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error(), GetLine())
		return

	}
	IdeaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	err = idea.CreateComment(userid, req.Content, IdeaId)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	var i model.IdeaModel
	model.DB.Self.Where("idea_id = ? ", IdeaId).First(&i)
	i.CommentSum++
	if err := model.DB.Self.Model(&i).Where("idea_id = ? ", IdeaId).Update("comment_sum", i.CommentSum).Error; err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	if err := message.CreateMessage(req.CommentedId, userid, 1, 0, IdeaId, i.Content, req.Content); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, errno.OK, nil)

}

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
// @Param token header string true  "userid"
// @Param req body idea.CreateCommentRequest true  "Add comment records to the database"
// @Param idea_id path int true "IdeaId"
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/detail/:id/comment/:idea_id [post]
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
	i2 := i
	i2.CommentSum++
	model.DB.Self.Model(&i).Update(i2)

	if err := message.CreateMessage(req.CommentedId, userid, 1, 0, IdeaId, i.Content, req.Content); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, errno.OK, nil)

}

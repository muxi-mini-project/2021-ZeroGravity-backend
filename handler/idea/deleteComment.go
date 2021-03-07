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

// @Summary Delete comment   
// @Description Delete the comment record from the database 
// @Tags idea
// @Accept  json
// @Produce  json
// @Param token header string true  "Delete the comment record from the database"
// @Success 200 "成功"
// @Router/api/v1/idea/detail/:id/comment/:comment_id [delete]
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
	// 调用服务
	if err := model.DeleteComment(id, uid); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	var i model.IdeaInfo
	model.DB.Self.Where("idea_id = ? ", id).First(&i)
	i2 := i
	i2.CommentSum--
	model.DB.Self.Model(&i).Update(i2)
	SendResponse(c, errno.OK, nil)

}

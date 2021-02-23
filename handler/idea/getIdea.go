package idea

import (
	"strconv"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/2021-ZeroGravity-backend/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetIdea ... 获取单个想法
func GetIdea(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 id 后面需要根据这个 id 判断是否为私密能获取
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	uid := c.MustGet("userID").(int)

	// 调用服务
	item, err := idea.GetIdea(id, uid)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &IdeaListItem{
		Id:          item.Id,
		Content:     item.Content,
		ReleaseDate: item.ReleaseDate,
		LikesSum:    item.LikesSum,
		CommentSum:  item.CommentSum,
		UserId:      item.UserId,
		Avatar:      item.Avatar,
		NickName:    item.NickName,
		Liked:       item.Liked,
	}

	SendResponse(c, errno.OK, resp)
}

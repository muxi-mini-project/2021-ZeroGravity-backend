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

// GetIdeaLike ... 获取某个用户点赞的想法列表
func GetIdeaLike(c *gin.Context) {
	log.Info("Idea getIdeaLike function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var id, limit, page int
	var err error

	// 从 token 获取 userid
	id = c.MustGet("userID").(int)

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	item, err := idea.GetIdeaLike(id, page*limit, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &IdeaResponse{}
	resp.Count = len(item)
	for _, v := range item {
		resp.List = append(resp.List, &IdeaListItem{
			Id:          v.Id,
			Content:     v.Content,
			ReleaseDate: v.ReleaseDate,
			LikesSum:    v.LikesSum,
			CommentSum:  v.CommentSum,
			UserId:      v.UserId,
			Avatar:      v.Avatar,
			NickName:    v.NickName,
			Liked:       v.Liked,
		})
	}

	SendResponse(c, errno.OK, resp)
}

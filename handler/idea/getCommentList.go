package idea

import (
	"strconv"

	"github.com/2021-ZeroGravity-backend/service/idea"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetCommentList ... 获取某个想法的评论列表
func GetCommentList(c *gin.Context) {
	log.Info("idea getComment function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var id, uid, limit, page int
	var err error

	// 从 token 获取 userid
	uid = c.MustGet("userID").(int)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

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
	item, err := idea.GetCommentList(id, uid, page*limit, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &CommentResponse{}
	resp.Count = len(item)
	for _, v := range item {
		resp.List = append(resp.List, &CommentListItem{
			Id:          v.Id,
			CommentedId: v.CommentedId,
			LikesSum:    v.LikesSum,
			ReleaseDate: v.ReleaseDate,
			Content:     v.Content,
			UserId:      v.UserId,
			Avatar:      v.Avatar,
			NickName:    v.NickName,
			Liked:       v.Liked,
		})
	}

	SendResponse(c, errno.OK, resp)
}

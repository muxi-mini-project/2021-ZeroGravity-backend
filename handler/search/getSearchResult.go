package search

import (
	"strconv"

	"github.com/2021-ZeroGravity-backend/service/search"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetSearchResult ... 获取搜索结果，结果为用户列表或想法列表
func GetSearchResult(c *gin.Context) {
	log.Info("Search getSearchResult function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取路由参数
	var page, limit, target int
	var keyword string
	var err error

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

	// 搜索目标，0->idea 1->user 默认搜索 idea
	target, err = strconv.Atoi(c.DefaultQuery("target", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	keyword = c.DefaultQuery("page", "")

	// 调用服务
	if target == 0 {
		// 搜索想法
		// token 获取 uid
		uid := c.MustGet("userID").(int)

		item, err := search.GetSearchIdeaResult(uid, page*limit, limit, keyword)
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
		return

	} else {
		// 搜索用户
		item, err := search.GetSearchUserResult(page*limit, limit, keyword)
		if err != nil {
			SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
			return
		}

		resp := &UserListResponse{}
		resp.Count = len(item)
		for _, v := range item {
			resp.List = append(resp.List, &UserInfo{
				Id:       v.Id,
				Avatar:   v.Avatar,
				NickName: v.NickName,
				Energy:   v.Energy,
			})
		}

		SendResponse(c, errno.OK, resp)
		return
	}
}

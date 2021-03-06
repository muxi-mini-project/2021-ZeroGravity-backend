package search

import (
	"strconv"

	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/service/search"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 获取搜索结果
// @Description 获取搜索结果，结果为用户列表或想法列表
// @Tags search
// @Accept  json
// @Produce  json
// @Param page  query int true "page--限制指定要检索的记录数"
// @Param limit query int true "limit--偏移量指定开始返回记录之前要跳过的记录数"
// @Param target query int true "target为0模糊搜索想法  target为1模糊搜索用户 若想模糊查找用户请给target赋值为1传给服务器"
// @Param  keyword query string true "关键字"
// @Param  id path int true  "userId"
// @Param  token header string true  "后端给前端的token"
// @Success 200 {object}  search.IdeaResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Success 200 {object}  search.UserListResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/search [get]
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

	keyword = c.DefaultQuery("keyword", "")
	//
	//
	//
	//获取用户id以创建搜索历史
	id := c.Param("id")
	model.CreateHistory(id, keyword)
	//
	//
	//
	// 调用服务
	if target == 0 {
		// 搜索想法
		// token 获取 uid
		uid := c.MustGet("userID").(int)

		item, err := search.GetSearchIdeaResult(uid, page*limit, limit, keyword)
		if err != nil {
			SendError(c, errno.ErrSearch, nil, err.Error(), GetLine())
			return
		}

		resp := &IdeaResponse{}
		resp.Count = len(item)
		resp.List = item
		SendResponse(c, errno.OK, resp)
	} else {
		// 搜索用户
		item, err := search.GetSearchUserResult(page*limit, limit, keyword)
		if err != nil {
			SendError(c, errno.ErrSearch, nil, err.Error(), GetLine())
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
	}
}

// @Summary 删除查询历史
// @Description 删除查询历史
// @Tags search
// @Accept  json
// @Produce  json
// @Param  id path int true  "userId--用户ID"
// @Param  h body model.History true  "delete history删除历史记录"
// @Success 200 "成功"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/search [delete]
//  DeleteHistory... 删除历史记录
func DeleteHistory(c *gin.Context) {
	log.Info("Message getMessageForComment function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	var h model.History
	c.BindJSON(&h)
	err := model.DeleteHistory(h)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, errno.OK, nil)
}

// @Summary 获取历史记录
// @Description 获取历史记录
// @Tags search
// @Accept  json
// @Produce  json
// @Param  id path int true  "userId--用户ID"
// @Success 200 {object}  []model.History  "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/search/histories [get]
//  GetHistories... 获取历史记录
func GetHistories(c *gin.Context) {
	log.Info("Message getMessageForComment function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))
	id := c.Param("id")
	histories, err := model.GetHistories(id)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}
	SendResponse(c, errno.OK, histories)
}

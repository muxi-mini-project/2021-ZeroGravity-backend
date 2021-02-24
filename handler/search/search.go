package search

import (
	"github.com/2021-ZeroGravity-backend/model"
)

// IdeaResponse ... 想法列表
type IdeaResponse struct {
	Count int                   `json:"count"`
	List  []*model.IdeaListItem `json:"list"`
}

// UserInfo ... 用户信息
type UserInfo struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
	Energy   int    `json:"energy"`
}

// UserListResponse ... 用户列表
type UserListResponse struct {
	Count int         `json:"count"`
	List  []*UserInfo `json:"list"`
}

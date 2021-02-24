package collection

import (
	"github.com/2021-ZeroGravity-backend/model"
)

// IdeaResponse ... 想法列表
type IdeaResponse struct {
	Count int                   `json:"count"`
	List  []*model.IdeaListItem `json:"list"`
}

// CreateCollectionRequest ... 新增收藏请求
type CreateCollectionRequest struct {
	CollectorId int `json:"collector_id"`
	IdeaId      int `json:"idea_id"`
}

// DeleteCollectionRequest ... 取消收藏请求
type DeleteCollectionRequest struct {
	CollectorId int `json:"collector_id"`
	IdeaId      int `json:"idea_id"`
}

package idea

import "github.com/2021-ZeroGravity-backend/model"

// CreateCommentRequest ... 创建评论请求
// 需要传递 被评论的 idea 的 userId
type CreateCommentRequest struct {
	CommentedId int    `json:"commented_id"`
	Content     string `json:"content"`
}

// DeleteCommentRequest ... 删除评论请求
type DeleteCommentRequest struct {
	Id int `json:"id"`
}

// CreateIdeaRequest ... 创建想法请求
type CreateIdeaRequest struct {
	Content string `json:"content"`
}

// DeleteIdeaRequest ... 删除想法请求
type DeleteIdeaRequest struct {
	IdeaId int `json:"idea_id"`
}

// UpdateIdeaLikeRequest ... 想法点赞/取消请求
type UpdateLikeRequest struct {
	Choice int `json:"choice"`
}

// IdeaResponse ... 想法列表
type IdeaResponse struct {
	Count int                   `json:"count"`
	List  []*model.IdeaListItem `json:"list"`
}

// CommentResponse ... 评论列表
type CommentResponse struct {
	Count int                      `json:"count"`
	List  []*model.CommentListItem `json:"list"`
}

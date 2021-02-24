package idea

import "github.com/2021-ZeroGravity-backend/model"

// CreateCommentRequest ... 创建评论请求
type CreateCommentRequest struct {
	CommenterId int    `json:"commenter_id"`
	CommentedId int    `json:"commented_id"`
	Content     string `json:"content"`
}

// DeleteCommentRequest ... 删除评论请求
type DeleteCommentRequest struct {
	CommenterId int `json:"commenter_id"`
	Id          int `json:"id"`
}

// UpdateCommentLikeRequest ... 评论点赞/取消请求
type UpdateCommentLikeRequest struct {
	CommentId int `json:"comment_id"`
	LikersId  int `json:"likers_id"`
	BelikedId int `json:"beliked_id"`
}

// CreateIdeaRequest ... 创建想法请求
type CreateIdeaRequest struct {
	PublisherId int    `json:"publisher_id"`
	Content     string `json:"content"`
	ReleaseDate string `json:"releaseDate"`
}

// DeleteIdeaRequest ... 删除想法请求
type DeleteIdeaRequest struct {
	PublisherId int `json:"publisher_id"`
	IdeaId      int `json:"idea_id"`
}

// UpdateIdeaLikeRequest ... 想法点赞/取消请求
type UpdateIdeaLikeRequest struct {
	IdeaId    int `json:"idea_id"`
	LikersId  int `json:"likers_id"`
	BelikedId int `json:"beliked_id"`
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

package report

// CreateReportRequest ... 新增举报请求
type CreateReportRequest struct {
	UserId    int    `json:"user_id"`
	Kind      int    `json:"kind"`
	Reason    string `json:"json"`
	CommentId int    `json:"comment_id"`
	IdeaId    int    `json:"idea_id"`
}

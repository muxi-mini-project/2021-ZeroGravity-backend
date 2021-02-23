package report

type CreateReportRequest struct {
	UserId    int    `json:"user_id"`
	Kind      int    `json:"kind"`
	Reason    string `json:"json"`
	CommentId int    `json:"comment_id"`
	IdeaId    int    `json:"idea_id"`
}

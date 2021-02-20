package idea

type CreateCommentRequest struct {
	CommenterId int    `json:"commenter_id"`
	CommentedId int    `json:"commented_id"`
	Content     string `json:"content"`
}

type DeleteCommentRequest struct {
	CommenterId int `json:"commenter_id"`
	Id          int `json:"id"`
}

type UpdateCommentLikeRequest struct {
	CommentId int `json:"comment_id"`
	LikersId  int `json:"likers_id"`
	BelikedId int `json:"beliked_id"`
}

type CreateIdeaRequest struct {
	PublisherId int    `json:"publisher_id"`
	Content     string `json:"content"`
	ReleaseDate string `json:"releaseDate"`
}

type DeleteIdeaRequest struct {
	PublisherId int `json:"publisher_id"`
	IdeaId      int `json:"idea_id"`
}

type UpdateIdeaLikeRequest struct {
	IdeaId    int `json:"idea_id"`
	LikersId  int `json:"likers_id"`
	BelikedId int `json:"beliked_id"`
}

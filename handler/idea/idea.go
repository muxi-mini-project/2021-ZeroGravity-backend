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

type CreateCollectionRequest struct {
	CollectorId int `json:"collector_id"`
	IdeaId      int `json:"idea_id"`
}
type DeleteCollectionRequest struct {
	CollectorId int `json:"collector_id"`
	IdeaId      int `json:"idea_id"`
}

type IdeaListItem struct {
	Id          int    `json:"idea_id"`
	Content     string `json:"content"`
	ReleaseDate string `json:"release_date"`
	LikesSum    int    `json:"likes_sum"`
	CommentSum  int    `json:"comment_sum"`
	UserId      int    `json:"publisher_id"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nickname"`
	Liked       bool   `json:"liked"` // 是否点赞
}

type IdeaResponse struct {
	Count int             `json:"count"`
	List  []*IdeaListItem `json:"list"`
}

type CommentListItem struct {
	Id          int    `json:"id"`
	CommentedId int    `json:"commented_id"`
	LikesSum    string `json:"likes_sum"`
	ReleaseDate string `json:"release_date"`
	Content     string `json:"content"`
	UserId      int    `json:"commenter_id"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nickname"`
	Liked       bool   `json:"liked"` // 是否点赞
}

type CommentResponse struct {
	Count int                `json:"count"`
	List  []*CommentListItem `json:"list"`
}

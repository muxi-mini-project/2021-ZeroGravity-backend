package collection

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

type CreateCollectionRequest struct {
	CollectorId int `json:"collector_id"`
	IdeaId      int `json:"idea_id"`
}
type DeleteCollectionRequest struct {
	CollectorId int `json:"collector_id"`
	IdeaId      int `json:"idea_id"`
}

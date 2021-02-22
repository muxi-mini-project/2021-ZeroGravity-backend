package search

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

type UserInfo struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
	Energy   int    `json:"energy"`
}

type UserListResponse struct {
	Count int         `json:"count"`
	List  []*UserInfo `json:"list"`
}

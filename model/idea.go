package model

type IdeaModel struct {
	IdeaId      int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
	Content     string `json:"content" gorm:"column:content;" binding:"required"`
	ReleaseDate string `json:"releaseDate" gorm:"column:releaseDate;" binding:"required"`
	PublisherId int    `json:"publisher_id" gorm:"column:publisher_id;" binding:"required"`
	LikesSum    int    `json:"likes_sum" gorm:"column:likes_sum;" binding:"required"`
	CommentSum  int    `json:"comment_sum" gorm:"column:comment_sum;" binding:"required"`
}

type IdeaInfo struct {
	Id          int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
	Content     string `json:"content" gorm:"column:content;" binding:"required"`
	ReleaseDate string `json:"release_date" gorm:"column:release_date;" binding:"required"`
	LikesSum    int    `json:"likes_sum" gorm:"column:likes_sum;" binding:"required"`
	CommentSum  int    `json:"comment_sum" gorm:"column:comment_sum;" binding:"required"`
	UserId      int    `json:"publisher_id" gorm:"column:publisher_id;" binding:"required"`
	Avatar      string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	NickName    string `json:"nickname" gorm:"column:nickname;" binding:"required"`
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

func (u *IdeaModel) TableName() string {
	return "tbl_idea"
}

func (u *IdeaModel) Create() error {
	return DB.Self.Create(u).Error
}

func DeleteIdea(id, uid int) error {
	u := &IdeaModel{}
	u.IdeaId = id
	d := DB.Self.Where("publisher_id = ?", uid).Delete(u)
	return d.Error
}

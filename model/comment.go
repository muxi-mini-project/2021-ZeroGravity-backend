package model

type CommentModel struct {
	Id          int    `json:"id" gorm:"column:id;" binding:"required"`
	CommenterId int    `json:"commenter_id" gorm:"column:commenter_id;" binding:"required"`
	CommentedId int    `json:"commented_id" gorm:"column:commented_id;" binding:"required"`
	LikesSum    string `json:"likes_sum" gorm:"column:likes_sum;" binding:"required"`
	ReleaseDate string `json:"release_date" gorm:"column:release_date;" binding:"required"`
	Content     string `json:"content" gorm:"column:content;" binding:"required"`
}

type CommentInfo struct {
	Id          int    `json:"id" gorm:"column:id;" binding:"required"`
	CommentedId int    `json:"commented_id" gorm:"column:commented_id;" binding:"required"`
	LikesSum    string `json:"likes_sum" gorm:"column:likes_sum;" binding:"required"`
	ReleaseDate string `json:"release_date" gorm:"column:release_date;" binding:"required"`
	Content     string `json:"content" gorm:"column:content;" binding:"required"`
	UserId      int    `json:"commenter_id" gorm:"column:commenter_id;" binding:"required"`
	Avatar      string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	NickName    string `json:"nickname" gorm:"column:nickname;" binding:"required"`
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

func (u *CommentModel) TableName() string {
	return "tbl_comments"
}

func (u *CommentModel) Create() error {
	return DB.Self.Create(u).Error
}
func DeleteComment(id, uid int) error {
	u := &CommentModel{}
	u.Id = id
	d := DB.Self.Where("commenter_id = ?", uid).Delete(u)
	return d.Error
}

func GetCommentListByIdeaId(id, offset, limit int) ([]*CommentInfo, []int, error) {
	var commentList = make([]*CommentInfo, 0)

	query := DB.Self.Table("tbl_comments").
		Select("tbl_comments.*,tbl_user.nickname,tbl_user.avatar").
		Where("tbl_comments.commeted_id = ?", id).
		Joins("left join tbl_user on tbl_user.id = tbl_comments.commenter_id").
		Order("tbl_comments.id desc")

	if err := query.Scan(&commentList).Error; err != nil {
		return nil, nil, err
	}

	var idList []int
	for _, v := range commentList {
		idList = append(idList, v.Id)
	}

	return commentList, idList, nil
}

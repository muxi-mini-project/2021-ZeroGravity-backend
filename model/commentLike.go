package model

type CommentLikeModel struct {
	Id        int `json:"id" gorm:"column:id;" binding:"required"`
	CommentId int `json:"comment_id" gorm:"comment_id;" binding:"required"`
	LikersId  int `json:"likers_id" gorm:"likers_id;" binding:"required"`
	BelikedId int `json:"beliked_id" gorm:"beliked_id;" binding:"required"`
}

func (u *CommentLikeModel) TableName() string {
	return "tbl_like_record_comment"
}

func (u *CommentLikeModel) Create() error {
	return DB.Self.Create(u).Error
}

func DeleteCommentLike(id int, uid string) error {
	u := &CommentLikeModel{}
	u.Id = id
	d := DB.Self.Where("likers_id = ?", uid).Delete(u)
	return d.Error
}

func GetCommentLikeForUser(uid int, scope []int) ([]*CommentLikeModel, error) {
	u := make([]*CommentLikeModel, 0)

	d := DB.Self.Table("tbl_like_record_comment").
		Where("likers_id = ? AND comment_id in (?)", uid, scope).
		Order("comment_id desc").Scan(&u)

	return u, d.Error
}

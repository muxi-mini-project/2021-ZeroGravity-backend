package model

type CommentLikeModel struct {
	Id        int    `json:"id" gorm:"column:id;" binding:"required"`
	CommentId int    `json:"comment_id" gorm:"comment_id;" binding:"required"`
	LikersId  string `json:"likers_id" gorm:"likers_id;" binding:"required"`
	BelikedId string `json:"beliked_id" gorm:"beliked_id;" binding:"required"`
}

func (u *CommentLikeModel) TableName() string {
	return "tbl_like_record_comment"
}

func (u *CommentLikeModel) Create() error {
	return DB.Self.Create(u).Error
}

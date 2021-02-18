package model

type IdeaLikeModel struct {
	Id        int    `json:"id" gorm:"column:id;" binding:"required"`
	IdeaId    int    `json:"idea_id" gorm:"idea_id;" binding:"required"`
	LikersId  string `json:"likers_id" gorm:"likers_id;" binding:"required"`
	BelikedId string `json:"beliked_id" gorm:"beliked_id;" binding:"required"`
}

func (u *IdeaLikeModel) TableName() string {
	return "tbl_like_record_idea"
}

func (u *IdeaLikeModel) Create() error {
	return DB.Self.Create(u).Error
}

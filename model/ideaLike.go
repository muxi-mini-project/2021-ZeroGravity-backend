package model

type IdeaLikeModel struct {
	Id        int `json:"id" gorm:"column:id;" binding:"required"`
	IdeaId    int `json:"idea_id" gorm:"idea_id;" binding:"required"`
	LikersId  int `json:"likers_id" gorm:"likers_id;" binding:"required"`
	BelikedId int `json:"beliked_id" gorm:"beliked_id;" binding:"required"`
}

func (u *IdeaLikeModel) TableName() string {
	return "tbl_like_record_idea"
}

func (u *IdeaLikeModel) Create() error {
	return DB.Self.Create(u).Error
}

func DeleteIdeaLike(id int, uid string) error {
	u := &IdeaLikeModel{}
	u.Id = id
	d := DB.Self.Where("likers_id = ?", uid).Delete(u)
	return d.Error
}

func GetIdeaLikeRecordForUser(id int, scope []int) ([]*IdeaLikeModel, error) {
	ideaLikeList := make([]*IdeaLikeModel, 0)
	d := DB.Self.Table("tbl_like_record_idea").Where("likers_id = ? AND idea_id in (?)", id, scope).Order("idea_id desc").Scan(&ideaLikeList)
	return ideaLikeList, d.Error
}

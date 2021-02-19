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

func DeleteIdeaLike(id int , uid string) error {
	u := &IdeaLikeModel{}
	u.Id = id
	d := DB.Self.Where("likers_id = ?", uid).Delete(u)
	return d.Error
}
type CreateIdeaLikeRequest struct {
	IdeaId    int    `json:"idea_id"`
	LikersId  string `json:"likers_id"`
	BelikedId string `json:"beliked_id"`
	
}


type DeleteIdeaLikeRequest struct {
	IdeaId    int    `json:"idea_id"`
	LikersId  string `json:"likers_id"`
	
	
}
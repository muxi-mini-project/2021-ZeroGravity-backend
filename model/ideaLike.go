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

func GetIdeaLikeByUserId(id, offset, limit int) ([]*IdeaInfo, error) {
	item := make([]*IdeaLikeModel, 0)

	d := DB.Self.Table("tbl_Like_record_idea").
		Where("likers_id  = ?", id).
		Offset(offset).Limit(limit).
		Order("id desc").Scan(&item)

	if d.Error != nil {
		return nil, d.Error
	}

	var idList []int
	for _, v := range item {
		idList = append(idList, v.IdeaId)
	}

	var ideaList = make([]*IdeaInfo, 0)

	query := DB.Self.Table("tbl_idea").
		Select("tbl_idea.*,tbl_user.nickname,tbl_user.avatar").
		Where("tbl_idea.idea_id IN ?", idList).
		Joins("left join tbl_user on tbl_user.id = tbl_idea.publisher_id").
		Order("tbl_idea.id desc")

	if err := query.Scan(&ideaList).Error; err != nil {
		return nil, err
	}

	return ideaList, nil
}

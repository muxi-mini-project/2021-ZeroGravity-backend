package model

type CollectionModel struct {
	Id          int `json:"id" gorm:"column:id;" binding:"required"`
	CollectorId int `json:"collector_id" gorm:"collector_id;" binding:"required"`
	IdeaId      int `json:"idea_id" gorm:"idea_id" binding:"required"`
}

func (u *CollectionModel) TableName() string {
	return "tbl_favorite_records"
}

func (u *CollectionModel) Create() error {
	return DB.Self.Create(u).Error
}

func DeleteCollection(id, uid int) error {
	u := &CollectionModel{}
	d := DB.Self.Where("collector_id = ? AND idea_id = ?", uid, id).Delete(u)
	return d.Error
}

// GetCollectionByUserId ... 通过 id 获取收藏
func GetCollectionByUserId(id, offset, limit int) ([]*IdeaInfo, []int, error) {
	item := make([]*CollectionModel, 0)

	d := DB.Self.Table("tbl_favorite_records").
		Where("collector_id = ?", id).
		Offset(offset).Limit(limit).
		Order("id desc").Scan(&item)

	if d.Error != nil {
		return nil, nil, d.Error
	}

	var idList []int
	for _, v := range item {
		idList = append(idList, v.IdeaId)
	}

	var ideaList = make([]*IdeaInfo, 0)

	query := DB.Self.Table("tbl_idea").
		Select("tbl_idea.*,tbl_user.nickname,tbl_user.avatar").
		Where("tbl_idea.idea_id IN (?)", idList).
		Joins("left join tbl_user on tbl_user.id = tbl_idea.publisher_id").
		Order("tbl_idea.idea_id desc")

	if err := query.Scan(&ideaList).Error; err != nil {
		return nil, nil, err
	}

	return ideaList, idList, nil
}

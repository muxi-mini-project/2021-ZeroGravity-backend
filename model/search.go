package model

func AgainstAndMatchIdea(offset, limit int, kw string) ([]*IdeaInfo, error) {
	ideaList := make([]*IdeaInfo, 0)

	// 搜索语句
	where := "MATCH (tbl_idea.content) AGAINST ('" + kw + "')"

	query := DB.Self.Debug().Table("tbl_idea").
		Select("tbl_idea.*,tbl_user.nickname,tbl_user.avatar").
		Joins("left join tbl_user on tbl_user.id = tbl_idea.publisher_id").
		Offset(offset).Limit(limit).Order("tbl_idea.idea_id desc").
		Where(where)

	if err := query.Scan(&ideaList).Error; err != nil {
		return nil, err
	}

	return ideaList, nil
}

func AgainstAndMatchUser(offset, limit int, kw string) ([]*UserModel, error) {
	userList := make([]*UserModel, 0)

	// 搜索语句
	where := "MATCH (nickname) AGAINST ('" + kw + "')"

	query := DB.Self.Debug().Table("tbl_user").
		Offset(offset).Limit(limit).Order("tbl_idea.idea_id desc").
		Where(where)

	if err := query.Scan(&userList).Error; err != nil {
		return nil, err
	}

	return userList, nil
}

func CreateHistory(id string, name string) error {
	var h History
	h.Name = name
	h.UserID = id
	if result := DB.Self.Create(&h); result.Error != nil {
		return result.Error
	}
	return nil
}
func GetHistories(id string) ([]History, error) {
	var histories []History
	if result := DB.Self.Where("user_id = ? ", id).Find(&histories); result.Error != nil {
		return histories, result.Error
	}
	return histories, nil
}
func DeleteHistory(h History) error {
	if result := DB.Self.Where("user_id = ? AND name = ? ", h.UserID, h.Name).Delete(&h); result.Error != nil {
		return result.Error
	}
	return nil
}

type History struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

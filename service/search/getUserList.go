package search

import (
	"github.com/2021-ZeroGravity-backend/model"
)

// GetSearchUserList ... 根据单字模糊搜索用户
func GetSearchUserList(offset, limit int, SingleKeyWord string) ([]*model.UserModel, error) {
	list, err := model.AgainstAndMatchUser(offset, limit, SingleKeyWord)
	if err != nil {
		return nil, err
	}

	return list, nil
}
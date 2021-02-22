package search

import (
	"github.com/2021-ZeroGravity-backend/model"
)

// GetSearchUserResult ... 搜索用户
func GetSearchUserResult(offset, limit int, keyword string) ([]*model.UserModel, error) {
	list, err := model.AgainstAndMatchUser(offset, limit, keyword)
	if err != nil {
		return nil, err
	}

	return list, nil
}

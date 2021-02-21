package user

import (
	"github.com/2021-ZeroGravity-backend/model"
)

func GetUserInfo(uid int) (*model.UserModel, error) {
	item, err := model.GetUserById(uid)
	if err != nil {
		return nil, err
	}

	return item, nil
}

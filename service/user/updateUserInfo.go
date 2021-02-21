package user

import (
	"github.com/2021-ZeroGravity-backend/model"
)

func UpdateUserInfo(uid int, avatar, nickname string) error {
	item, err := model.GetUserById(uid)
	if err != nil {
		return err
	}

	item.Avatar = avatar
	item.NickName = nickname

	return item.Update()
}

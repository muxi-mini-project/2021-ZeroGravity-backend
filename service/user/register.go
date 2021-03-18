
package user

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/2021-ZeroGravity-backend/model"
)

func Register(account, accountPassword, nickName, avatar string) error {
	if account == "" || accountPassword == "" || nickName == "" {
		return errors.New("lack for information")
	}

	// 判断用户是否存在
	user, err := model.GetUserByAccount(account)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &model.UserModel{
			Account:         account,
			AccountPassword: accountPassword,
			NickName:        nickName,
			Avatar:          avatar,
		}
	} else if err == nil {
		return errors.New("user had existed")
	} else {
		return err
	}

	return user.Create()
}
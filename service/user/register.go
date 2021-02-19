package service

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/2021-ZeroGravity-backend/model"
)

func Register(req *model.CreateUserRequest) error {
	if req.Account == "" || req.AccountPassword == "" || req.NickName == "" {
		return errors.New("lack for information")
	}

	// 判断用户是否存在
	user, err := model.GetUserByAccount(req.Account)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &model.UserModel{
			Account:         req.Account,
			AccountPassword: req.AccountPassword,
			NickName:        req.NickName,
			Avatar:          req.Avatar,
		}
	} else if err == nil {
		return errors.New("user had existed")
	} else {
		return err
	}

	return user.Create()
}

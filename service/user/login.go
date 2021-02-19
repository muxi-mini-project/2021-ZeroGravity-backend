package service

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/pkg/token"
	"github.com/2021-ZeroGravity-backend/util"
)

func Login(req *model.LoginRequest) (string, error) {
	user, err := model.GetUserByAccountAndPassword(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user doesn't exist")
		}
		return "", err
	}

	// 生成 auth token
	token, err := token.GenerateToken(&token.TokenPayload{
		ID:      user.Id,
		Expired: util.GetExpiredTime(),
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

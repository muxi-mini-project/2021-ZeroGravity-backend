package message

import (
	"github.com/2021-ZeroGravity-backend/model"
)

func UpdateMessageReadAll(uid int) error {
	return model.ReadAll(uid)
}

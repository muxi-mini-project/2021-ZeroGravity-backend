package message

import (
	"github.com/2021-ZeroGravity-backend/model"
)

func GetMessageTip(uid int) (bool, error) {
	count, err := model.GetMessageTipByUserId(uid)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

package user

import (
	"fmt"

	"github.com/2021-ZeroGravity-backend/model"
)

//DeleteCollection  is used to delete collection
func DeleteCollection(id, uid int) error {
	if err := model.DeleteCollion(id, uid); err != nil {
		fmt.Println(err)
	}

	return nil
}

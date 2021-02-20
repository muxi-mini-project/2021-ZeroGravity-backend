package idea

import (
	"fmt"

	"github.com/2021-ZeroGravity-backend/model"
)

//DeleteComment   is used to delete comment
func DeleteComment(id, uid int) error {
	if err := model.DeleteComment(id, uid); err != nil {
		fmt.Println(err)
	}

	return nil

}

package idea

import (
	"fmt"

	"github.com/2021-ZeroGravity-backend/model"
)

// DeleteIdea is used to delete idea
func DeleteIdea(id, uid int) error {
	if err := model.DeleteIdea(id, uid); err != nil {

		fmt.Println(err)
	}

	return nil
}

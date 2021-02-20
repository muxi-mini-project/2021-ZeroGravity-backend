package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
)

//CreateComment is used to create comment
func CreateComment(id, uid int, content string) error {

	var comment *model.CommentModel

	comment = &model.CommentModel{

		CommenterId: uid,
		CommentedId: id,
		Content:     content,
	}

	return comment.Create()

}

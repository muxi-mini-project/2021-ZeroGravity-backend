package idea

import (
	"time"

	"github.com/2021-ZeroGravity-backend/model"
)

//CreateComment is used to create comment
func CreateComment(id, uid int, content string, IdeaID int) error {

	var comment model.CommentModel

	comment = model.CommentModel{
		LikesSum:    0,
		CommenterId: uid,
		CommentedId: id,
		Content:     content,
		ReleaseDate: time.Now().Format("2006-01-02 15:04:05"),
		IdeaId:      IdeaID,
	}
	return comment.Create()
}

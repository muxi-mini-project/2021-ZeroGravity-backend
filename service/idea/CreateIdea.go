package idea

import (
	"time"

	"github.com/2021-ZeroGravity-backend/model"
)

//CreateIdea is used to create idea
func CreateIdea(content string, uid int) error {

	var idea *model.IdeaModel

	t := time.Now()

	idea = &model.IdeaModel{
		PublisherId: uid,
		Content:     content,
		ReleaseDate: t.Format("2006-01-02 15:04:05"),
	}

	return idea.Create()

}

package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
)

//CreateIdea is used to create idea
func CreateIdea(content string, uid int, date string) error {

	var idea *model.IdeaModel

	idea = &model.IdeaModel{
		PublisherId: uid,
		Content:     content,
		ReleaseDate: date,
	}

	return idea.Create()

}

package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
)

//UpdateIdeaLike is used to update idea like
func UpdateIdeaLike(id, uid int) error {

	var IdeaLike *model.IdeaLikeModel

	IdeaLike = &model.IdeaLikeModel{
		IdeaId:   id,
		LikersId: uid,
	}

	return IdeaLike.Create()

}

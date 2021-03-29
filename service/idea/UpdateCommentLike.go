package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
)

//UpdateCommentLike is used to update comment like
func UpdateCommentLike(id, uid int) error {

	var CommentLike *model.CommentLikeModel

	CommentLike = &model.CommentLikeModel{
		CommentId: id,
		LikersId:  uid,
		//BelikedId: uid2,
	}

	return CommentLike.Create()

}

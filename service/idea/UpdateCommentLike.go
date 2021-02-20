package idea

import (

"github.com/2021-ZeroGravity-backend/model"

)

//UpdateCommentLike is used to update comment like
    func UpdateCommentLike (req *model.CreateCommentLikeRequest) error {
	
	    var CommentLike   *model.CommentLikeModel 

	    CommentLike  = &model.CommentLikeModel {
  
          CommentId:              req.CommentId,
	      LikersId:               req.LikersId,
	      BelikedId:              req.BelikedId,
	                           

    }

    return  CommentLike.Create()

}
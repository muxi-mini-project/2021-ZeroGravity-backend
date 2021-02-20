package idea

import (

"github.com/2021-ZeroGravity-backend/model"

)

//UpdateIdeaLike is used to update idea like
func UpdateIdeaLike (req *model.CreateIdeaLikeRequest) error {
	
	var IdeaLike   *model.IdeaLikeModel 

	IdeaLike  = &model.IdeaLikeModel {

	  IdeaId:                 req.IdeaId,
	  LikersId:               req.LikersId,
	  BelikedId:              req.BelikedId,
						   

}

    return  IdeaLike.Create()

}
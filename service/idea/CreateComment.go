package service

import (

"github.com/2021-ZeroGravity-backend/model"

)
//CreateComment is used to create comment
    func CreateComment (req *model.CreateCommentRequest) error {
	
	   var comment   *model.CommentModel
	
	   comment = &model.CommentModel{
		
		CommenterId:         req.CommenterId,
		CommentedId:         req.CommentedId,
		Content:             req.Content,
	
    }

	return comment.Create()

}
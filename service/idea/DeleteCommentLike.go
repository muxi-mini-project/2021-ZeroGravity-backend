package service

import (
   "fmt"
   "github.com/2021-ZeroGravity-backend/model"

)

//DeleteCommentLike   is used to delete comment like
    func DeleteCommentLike (req *model.DeleteCommentRequest) error {

	  uid := req.CommenterId;
	  id  := req.Id;
	  err := model.DeleteCommentLike(id, uid )
	  if err != nil {
	
	  fmt.Println(err)
    }

    return nil 

}
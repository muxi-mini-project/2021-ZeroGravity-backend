package service

import (
   "fmt"
   "github.com/2021-ZeroGravity-backend/model"

)

//DeleteComment   is used to delete comment
    func DeleteComment (req *model.DeleteCommentRequest) error {

	  uid := req.CommenterId;
	  id  := req.Id;
	  err := model.DeleteComment(id, uid )
	  if err != nil {
	
	  fmt.Println(err)
    }

    return nil 

}
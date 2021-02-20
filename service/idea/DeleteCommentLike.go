package idea

import (
   "fmt"
   "github.com/2021-ZeroGravity-backend/model"

)

//DeleteCommentLike   is used to delete comment like
    func DeleteCommentLike (req *model.DeleteCommentLikeRequest) error {

	  uid := req.LikersId;
	  id  := req.CommentId;
	  err := model.DeleteCommentLike(id, uid )
	  if err != nil {
	
	  fmt.Println(err)
    }

    return nil 

}
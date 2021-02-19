package service

import (
   "fmt"
   "github.com/2021-ZeroGravity-backend/model"

)
// DeleteIdea is used to delete idea
    func DeleteIdea (req *model.DeleteIdeaRequest) error {

	  uid := req.PublisherId;
	  id  := req.IdeaId;
	  err := model.DeleteIdea(id, uid )
	  if err != nil {
  
	  fmt.Println(err)
    }

  return nil 
}
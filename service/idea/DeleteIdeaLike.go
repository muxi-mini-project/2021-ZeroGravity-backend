package service
import (
   "fmt"
   "github.com/2021-ZeroGravity-backend/model"

)

//DeleteIdeaLike   is used to delete idea like
    func DeleteIdeaLike (req *model.DeleteIdeaRequest) error {

	  uid := req.PublisherId;
	  id  := req.IdeaId;
	  err := model.DeleteIdeaLike(id, uid )
	  if err != nil {
	
	  fmt.Println(err)
    }

    return nil 

}
package user

import (
   "fmt"
   "github.com/2021-ZeroGravity-backend/model"

)

//DeleteCollection  is used to delete collection
    func DeleteCollection (req *model.DeleteCollectionRequest) error {

	    uid := req.CollectorId;
		id  := req.IdeaId;
	    err := model.DeleteCollion(id, uid )
	    if err != nil {
		
		  fmt.Println(err)
	   }

	 return nil 
   
	}
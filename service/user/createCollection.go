package user

import (

"github.com/2021-ZeroGravity-backend/model"

)
//CreateCollection is used to create collection
   func CreateCollection (req *model.CreateCollectionRequest) error {
	
	   var collection   *model.CollectionModel 
 
       collection = &model.CollectionModel {
	 
	    CollectorId:         req.CollectorId,
	    IdeaId:              req.IdeaId,
	
 
    }

    return  collection.Create()

}
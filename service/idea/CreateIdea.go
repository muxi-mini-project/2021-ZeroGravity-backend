
package idea

import (
		
"github.com/2021-ZeroGravity-backend/model"

)


//CreateIdea is used to create idea
    func CreateIdea (req *model.CreateIdeaRequest) error {
	
	    var idea  *model.IdeaModel
	
	    idea = &model.IdeaModel{
		
		  PublisherId:         req.PublisherId,
		  Content:             req.Content,
		  ReleaseDate:         req.ReleaseDate,
	
    }

	return idea.Create()

}
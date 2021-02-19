package service

import (
	"fmt"

	




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
//CreateCollection is used to create collection
   func CreateCollection (req *model.CreateCollectionRequest) error {
	
	   var collection   *model.CollectionModel 
 
       collection = &model.CollectionModel {
	 
	    CollectorId:         req.CollectorId,
	    IdeaId:              req.IdeaId,
	
 
    }

    return  collection.Create()

}
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

//UpdateCommentLike is used to update comment like
    func UpdateCommentLike (req *model.CreateCommentLikeRequest) error {
	
	    var CommentLike   *model.CommentLikeModel 

	    CommentLike  = &model.CommentLikeModel {
  
          CommentId:              req.CommentId,
	      LikersId:               req.LikersId,
	      BelikedId:              req.BelikedId,
	                           

    }

    return  CommentLike.Create()

}
//DeleteCommentLike  is used to delete comment like
   func DeleteCommentLike (req *model.DeleteCommentLikeRequest) error {

	  uid := req.LikersId;
	  id  := req.CommentId;
	  err := model.DeleteCommentLike(id, uid )
	  if err != nil {
	
	  fmt.Println(err)
    }

    return nil 

}


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
//DeleteIdeaLike  is used to  delete idea like
func DeleteIdeaLike (req *model.DeleteIdeaLikeRequest) error {

	uid := req.LikersId;
	id  := req.IdeaId;
	err := model.DeleteIdeaLike(id,  uid )
	if err != nil {
  
	fmt.Println(err)
  }

  return nil 

}
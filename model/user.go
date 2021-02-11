package model

import(
	"github.com/2021-ZeroGravity-backend/model"

	"github.com/gin-gonic/gin"
	"log"
)


type User struct {
	Id             string    `json:"id"`
	Account        string    `json:"account"`
	Nickname       string    `json:"nickname"`
	Account        string    `json:"account"`
	Avatar         string    `json:"avatar"`
	Energy         int       `json:"energy"`
}


func NewUser(user User) string {
	
	if err := model.DB.Table("tbl_user").Create(&user); err != nil {
		log.Println("NewUser" ) 
	}
	return user.Nickname
}











type Idea struct{
	IdeaId         int       `json:"idea_id"`
    Content        string    `json:"content"`
	ReleaseDate	   string    `json:"releaseDate"`
	PublisherId    string 	 `json:"publisher_id"`
	likessum	   int 	     `json:"likes_sum"`
	CommentSum	   int       `json:"comment_sum"`
	CollectionSum  int 		 `json:"collection_sum"`
}
 


func NewIdea (idea Idea) string {
	
	if err := model.DB.Table("tbl_idea").Create(&idea); err != nil {
		log.Println("NewIdea")
		
	}
	return idea.Content
}


func DeleteIdea (idea Idea) string {
	const result = "删除成功"
	if err := model.DB.Table("tbl_idea").Delete((&idea),1); err != nil {
		log.Println("DeleteIdea " )
		
	}
	return   result

}



type Comment struct{
	Id             int       `json:"id"`
	CommenterId    string    `json:"commenter_id"`
	CommentedId    string    `json:"commented_id"`
	likessum       string    `json:"likes_sum"`
	ReleaseDate    string    `json:"release_date"`
	Content        string    `json:"content"`

}


func NewComment(comment  Comment) string {
	
	if err := model.DB.Table("tbl_comments").Create(&comment); err != nil {
		log.Println("NewComment")
		
	}
	return comment.Content 
}






type Collection struct{
	Id             int       `json:"id"`
	CollectorId    int       `json:"collector_id"`
	IdeaId         int       `json:"idea_id"`

}

func NewCollection(collection Collection) int {
	
	if err := model.DB.Table("tbl_favorite_records").Create(&collection); err != nil {
		log.Println("NewCollection")
		
	}
	return collection.Id
}






type IdeaLikeRecord  struct{
	Id             int       `json:"id"`
	IdeaId         int       `json:"idea_id"`
	LikersId       string    `json:"likers_id"`
	BelikedId      string    `json:"beliked_id"`
    
}

func NewIdeaLike(ideaLikeRecord IdeaLikeRecord) string {
	
	if err := model.DB.Table("tbl_like_record_idea").Create(&ideaLikeRecord); err != nil {
		log.Println("NewIdeaLike" )
	
	}
	return ideaLikeRecord.BelikedId
}




type CommentLikeRecord  struct{
	Id             int       `json:"id"`
	CommentId      int       `json:"comment_id"`
	LikersId       string    `json:"likers_id"`
	BelikedId      string    `json:"beliked_id"`
}


func NewCommentLike(commentLikeRecord CommentLikeRecord) string  {
	
	if err :=model.DB.Table("tbl_like_record_comment").Create(&commentLikeRecord); err != nil {
		log.Println("NewCommentLike" )
		
	}
	return commentLikeRecord.BelikedId
}










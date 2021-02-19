package model

type CommentModel struct {
	Id          int    `json:"id" gorm:"column:id;" binding:"required"`
	CommenterId string `json:"commenter_id" gorm:"column:commenter_id;" binding:"required"`
	CommentedId string `json:"commented_id" gorm:"column:commented_id;" binding:"required"`
	likesSum    string `json:"likes_sum" gorm:"column:likes_sum;" binding:"required"`
	ReleaseDate string `json:"release_date" gorm:"column:release_date;" binding:"required"`
	Content     string `json:"content" gorm:"column:content;" binding:"required"`
}

func (u *CommentModel) TableName() string {
	return "tbl_comments"
}

func (u *CommentModel) Create() error {
	return DB.Self.Create(u).Error
}
func DeleteComment(id int, uid string) error {
	u := &CommentModel{}
	u.Id = id
	d := DB.Self.Where("commenter_id = ?", uid).Delete(u)
	return d.Error
}
// --Request&Response--
 
type CreateCommentRequest struct {
	CommenterId string `json:"commenter_id`
	CommentedId string `json:"commented_id"`
	Content     string `json:"content"`

	
}

type DeleteCommentRequest struct {
	CommenterId string `json:"commenter_id`
	Id          int    `json:"id"`


}
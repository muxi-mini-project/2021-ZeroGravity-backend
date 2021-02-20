package model

type IdeaModel struct {
	IdeaId        int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
	Content       string `json:"content" gorm:"column:content;" binding:"required"`
	ReleaseDate   string `json:"releaseDate" gorm:"column:releaseDate;" binding:"required"`
	PublisherId   string `json:"publisher_id" gorm:"column:publisher_id;" binding:"required"`
	likessum      int    `json:"likes_sum" gorm:"column:likes_sum;" binding:"required"`
	CommentSum    int    `json:"comment_sum" gorm:"column:comment_sum;" binding:"required"`
	
}

func (u *IdeaModel) TableName() string {
	return "tbl_idea"
}

func (u *IdeaModel) Create() error {
	return DB.Self.Create(u).Error
}

func DeleteIdea(id int , uid string) error {
	u := &IdeaModel{}
	u.IdeaId = id
	d := DB.Self.Where("publisher_id = ?", uid).Delete(u)
	return d.Error
}


// --Request&Response--
type CreateIdeaRequest struct {
	PublisherId     string `json:"publisher_id"`
	Content         string `json:"content"`
	ReleaseDate     string `json:"releaseDate"`
	
}


type DeleteIdeaRequest struct {
	PublisherId     string  `json:"publisher_id"`
	IdeaId          int       `json:"idea_id"`
	
	
}
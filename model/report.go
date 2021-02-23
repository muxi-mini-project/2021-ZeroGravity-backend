package model

type ReportModel struct {
	Id        int    `json:"id" gorm:"column:id;" binding:"required"`
	UserId    int    `json:"user_id" gorm:"column:user_id;" binding:"required"`
	Reporter  int    `json:"reporter" gorm:"column:reporter;" binding:"required"`
	Kind      int    `json:"kind" gorm:"column:kind;" binding:"required"`
	Reason    string `json:"reason" gorm:"column:reason;" binding:"required"`
	CommentId int    `json:"comment_id" gorm:"column:comment_id;" binding:"requried"`
	IdeaId    int    `json:"idea_id" gorm:"column:idea_id;" binding:"requried"`
}

func (u *ReportModel) TableName() string {
	return "tbl_report"
}

func (u *ReportModel) Create() error {
	return DB.Self.Create(u).Error
}

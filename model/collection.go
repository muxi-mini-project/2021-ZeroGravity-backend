package model

type CollectionModel struct {
	Id          int `json:"id" gorm:"column:id;" binding:"required"`
	CollectorId int `json:"collector_id" gorm:"collector_id;" binding:"required"`
	IdeaId      int `json:"idea_id" gorm:"idea_id" binding:"required"`
}

func (u *CollectionModel) TableName() string {
	return "tbl_favorite_records"
}

func (u *CollectionModel) Create() error {
	return DB.Self.Create(u).Error
}

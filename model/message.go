package model

import (
	"github.com/2021-ZeroGravity-backend/pkg/constvar"
)

type MessageModel struct {
	Id        int    `json:"id" gorm:"column:id;" binding:"required"`
	PubUserId int    `json:"pub_user_id" gorm:"column:pub_user_id;" binding:"required"`
	SubUserId int    `json:"sub_user_id" gorm:"column:sub_user_id;" binding:"required"`
	Kind      int    `json:"kind" gorm:"column:kind;" binding:"required"`
	IsRead    int    `json:"is_read" gorm:"column:"is_read;" binding:"required""`
	Date      string `json:"date" gorm:"column:date;" binding:"required"`
	CommentId int    `json:"comment_id" gorm:"column:comment_id;" binding:"required"`
	Content   string `json:""`
	IdeaId    int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
}

func (u *MessageModel) TableName() string {
	return "tbl_message"
}

func (u *MessageModel) Create() error {
	return DB.Self.Create(u).Error
}

func GetMessageTipByUserId(uid int) (int, error) {
	var count int
	d := DB.Self.Table("tbl_message").Where("sub_user_id = ? AND is_read = ?", uid, 0).Count(&count)
	return count, d.Error
}

func ReadAll(uid int) error {
	d := DB.Self.Table("message").Where("sub_user_id = ?", uid).Update("is_read", 1)
	return d.Error
}

// GetMessageForComment ... 拉取评论我的 idList
func GetMessageForComment(uid, offset, limit int) ([]*MessageModel, error) {
	message := make([]*MessageModel, 0)

	query := DB.Self.Table("tbl_message").
		Where("sub_user_id = ? AND kind = ?", uid, constvar.CommentMessage).
		Offset(offset).Limit(limit)

	if err := query.Scan(&message).Error; err != nil {
		return nil, err
	}

	return message, nil
}

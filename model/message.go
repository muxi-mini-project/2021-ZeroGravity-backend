package model

import (
	"github.com/2021-ZeroGravity-backend/pkg/constvar"
)

type MessageModel struct {
	Id        int    `json:"id" gorm:"column:id;" binding:"required"`
	PubUserId int    `json:"pub_user_id" gorm:"column:pub_user_id;" binding:"required"`
	SubUserId int    `json:"sub_user_id" gorm:"column:sub_user_id;" binding:"required"`
	Kind      int    `json:"kind" gorm:"column:kind;" binding:"required"`
	IsRead    int    `json:"is_read" gorm:"column:is_read;" binding:"required""`
	Date      string `json:"date" gorm:"column:date;" binding:"required"`
	CommentId int    `json:"comment_id" gorm:"column:comment_id;" binding:"required"`
	Reply     string `json:"reply" gorm:"column:reply;" binding:"required"`
	Content   string `json:"content" gorm:"column:content;" binding:"required"`
	IdeaId    int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
}

type MessageForCommentItem struct {
	IdeaId   int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
	Content  string `json:"content" gorm:"column:content;" binding:"required"`
	Reply    string `json:"reply" gorm:"column:reply;" binding:"required"`
	UserId   int    `json:"user_id" gorm:"column:pub_user_id;" binding:"required"`
	Date     string `json:"date" gorm:"column:date;" binding:"requried"`
	Avatar   string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	NickName string `json:"nickname" gorm:"column:nickname;" binding:"required"`
}

type MessageForLikeItem struct {
	IdeaId    int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
	Content   string `json:"content" gorm:"column:content;" binding:"required"`
	CommentId int    `json:"comment_id" gorm:"column:comment_id;" binding:"required"`
	UserId    int    `json:"user_id" gorm:"column:pub_user_id;" binding:"required"`
	Date      string `json:"date" gorm:"column:date;" binding:"requried"`
	Avatar    string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	NickName  string `json:"nickname" gorm:"column:nickname;" binding:"required"`
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
	d := DB.Self.Table("tbl_message").Where("sub_user_id = ?", uid).Update("is_read", 1)
	return d.Error
}

// GetMessageForComment ... 拉取评论我的消息
func GetMessageForComment(uid, offset, limit int) ([]*MessageForCommentItem, error) {
	message := make([]*MessageForCommentItem, 0)

	query := DB.Self.Table("tbl_message").
		Select("tbl_message.idea_id,tbl_message.content,tbl_message.reply,tbl_message.pub_user_id,tbl_message.date,tbl_user.nickname,tbl_user.avatar").
		Where("sub_user_id = ? AND kind = ?", uid, constvar.CommentMessage).
		Joins("left join tbl_user on tbl_user.id = tbl_message.pub_user_id").
		Offset(offset).Limit(limit)

	if err := query.Scan(&message).Error; err != nil {
		return nil, err
	}

	return message, nil
}



// GetUserList ...拉取单字搜索的用户信息

















// GetMessageForLike ... 拉取点赞我的消息
func GetMessageForLike(uid, offset, limit int) ([]*MessageForLikeItem, error) {
	message := make([]*MessageForLikeItem, 0)

	query := DB.Self.Table("tbl_message").
		Select("tbl_message.idea_id,tbl_message.content,tbl_message.comment_id,tbl_message.pub_user_id,tbl_message.date,tbl_user.nickname,tbl_user.avatar").
		Where("sub_user_id = ? AND kind = ?", uid, constvar.CommentMessage).
		Joins("left join tbl_user on tbl_user.id = tbl_message.pub_user_id").
		Offset(offset).Limit(limit)

	if err := query.Scan(&message).Error; err != nil {
		return nil, err
	}

	return message, nil
}

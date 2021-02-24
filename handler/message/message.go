package message

import (
	"github.com/2021-ZeroGravity-backend/model"
)

// GetMessageTipResponse ... 获取消息提示
type GetMessageTipResponse struct {
	HaveMessage bool `json:"is_message"`
}

// GetMessageForCommentResponse ... 评论信息列表
type GetMessageForCommentResponse struct {
	Count int                            `json:"count"`
	List  []*model.MessageForCommentItem `json:"list"`
}

// GetMessageForLikeResponse ... 点赞信息列表
type GetMessageForLikeResponse struct {
	Count int                         `json:"count"`
	List  []*model.MessageForLikeItem `json:"list"`
}

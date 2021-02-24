package message

import (
	"time"

	"github.com/2021-ZeroGravity-backend/model"
)

// CreateMessage ... 别的服务调用，不单独列 api
// 点赞消息传 pub sub kind commentId ideaId content
// 评论消息传 pub sub kind ideaId content reply
func CreateMessage(pub, sub, kind, a1, a2 int, a3, a4 string) error {
	t := time.Now()

	var message *model.MessageModel
	if kind == 0 {
		message.CommentId = a1
		message.IdeaId = a2
		message.Content = a3
	} else if kind == 1 {
		message.IdeaId = a1
		message.Content = a3
		message.Reply = a4
	}
	message = &model.MessageModel{
		PubUserId: pub,
		SubUserId: sub,
		Kind:      kind,
		Date:      t.Format("2006-01-02 15:04:05"),
	}

	return message.Create()

}

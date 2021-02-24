package message

import (
	"time"

	"github.com/2021-ZeroGravity-backend/model"
)

// CreateMessage ... 别的服务调用，不单独列 api
// 点赞消息传 pub sub kind commentId ideaId content
// 评论消息传 pub sub kind ideaId content reply
func CreateMessage(pub, sub, kind, commentId, ideaId int, content, reply string) error {
	t := time.Now()

	var message *model.MessageModel

	message = &model.MessageModel{
		PubUserId: pub,
		SubUserId: sub,
		Kind:      kind,
		CommentId: commentId,
		Reply:     reply,
		IdeaId:    ideaId,
		Content:   content,
		Date:      t.Format("2006-01-02 15:04:05"),
	}

	return message.Create()

}

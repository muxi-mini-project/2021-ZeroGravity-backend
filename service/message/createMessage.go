package message

import (
	"time"

	"github.com/2021-ZeroGravity-backend/model"
)

// CreateMessage ... 别的服务调用，不单独列 api
func CreateMessage(pub, sub, kind, commentId, IdeaId int) error {
	t := time.Now()

	var message *model.MessageModel

	message = &model.MessageModel{
		PubUserId: pub,
		SubUserId: sub,
		Kind:      kind,
		CommentId: commentId,
		IdeaId:    IdeaId,
		Date:      t.Format("2006-01-02 15:04:05"),
	}

	if err := message.Create(); err != nil {
		return err
	}

	return nil
}

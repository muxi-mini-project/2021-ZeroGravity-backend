package report

import (
	"github.com/2021-ZeroGravity-backend/model"
)

func CreatReport(uid, userId, kind, commentId, ideaId int, reason string) error {

	report := &model.ReportModel{
		UserId:    userId,
		Reporter:  uid,
		Kind:      kind,
		Reason:    reason,
		CommentId: commentId,
		IdeaId:    ideaId,
	}

	if err := report.Create(); err != nil {
		return err
	}

	return nil
}

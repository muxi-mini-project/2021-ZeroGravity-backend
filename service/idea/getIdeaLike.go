package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
)

func GetIdeaLike(id, offset, limit int) ([]*model.IdeaListItem, error) {
	list, err := model.GetIdeaLikeByUserId(id, offset, limit)
	if err != nil {
		return nil, err
	}

	// 合并
	listItem := make([]*model.IdeaListItem, 0)
	for _, item := range list {
		listItem = append(listItem, &model.IdeaListItem{
			Id:          item.Id,
			Content:     item.Content,
			ReleaseDate: item.ReleaseDate,
			LikesSum:    item.LikesSum,
			CommentSum:  item.CommentSum,
			UserId:      item.UserId,
			Avatar:      item.Avatar,
			NickName:    item.NickName,
			Liked:       true,
		})
	}

	return listItem, nil
}

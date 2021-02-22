package search

import (
	"github.com/2021-ZeroGravity-backend/model"
)

// GetSearchIdeaResult ... 搜索想法
func GetSearchIdeaResult(uid, offset, limit int, keyword string) ([]*model.IdeaListItem, error) {
	list, err := model.AgainstAndMatchIdea(offset, limit, keyword)
	if err != nil {
		return nil, err
	}

	var idList []int
	for _, v := range list {
		idList = append(idList, v.Id)
	}

	// 查询点赞情况
	ideaLikeList, err := model.GetIdeaLikeRecordForUser(uid, idList)
	if err != nil {
		return nil, err
	}

	// 合并
	listItem := make([]*model.IdeaListItem, 0)
	for _, item := range list {
		for _, v := range ideaLikeList {
			if v.IdeaId == item.Id {
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
				break
			}
		}
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

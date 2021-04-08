package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
)

func GetCollection(id, offset, limit int) ([]*model.IdeaListItem, error) {
	list, idList, err := model.GetCollectionByUserId(id, offset, limit)
	if err != nil {
		return nil, err
	}

	ideaLikeList, err := model.GetIdeaLikeRecordForUser(id, idList)
	if err != nil {
		return nil, err
	}

	// 合并
	listItem := make([]*model.IdeaListItem, 0)
	lenLike := len(ideaLikeList)
	lenList := len(list)
	i := 0 // list 的指针
	j := 0 // like 的指针
	for i < lenList && j < lenLike {
		if ideaLikeList[j].IdeaId > list[i].Id { // 如果 likelist 当前 id 比 idealist id 大，说明不在范围内，直接跳过。
			j++ // 只有 j 索引往后移动
			continue
		}
		if ideaLikeList[j].IdeaId == list[i].Id { // 如果 likelist 当前 id 等于 idealist id ，是该用户点赞的记录，liked 设置 1。
			item := list[i]
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
			i++
			j++
			continue
		}
		if ideaLikeList[j].IdeaId < list[i].Id { // 如果 likelist 当前 id 小于 idealist id，该记录不是目标记录， liked 设置成 0。
			item := list[i]
			listItem = append(listItem, &model.IdeaListItem{
				Id:          item.Id,
				Content:     item.Content,
				ReleaseDate: item.ReleaseDate,
				LikesSum:    item.LikesSum,
				CommentSum:  item.CommentSum,
				UserId:      item.UserId,
				Avatar:      item.Avatar,
				NickName:    item.NickName,
				Liked:       false,
			})
			i++ // 索引 i 往后走，j 等待目标
			continue
		}
	}

	// 若 idealist 没走完，需要把剩下的 idea 插入结果
	if i < lenList {
		for i < lenList {
			item := list[i]
			listItem = append(listItem, &model.IdeaListItem{
				Id:          item.Id,
				Content:     item.Content,
				ReleaseDate: item.ReleaseDate,
				LikesSum:    item.LikesSum,
				CommentSum:  item.CommentSum,
				UserId:      item.UserId,
				Avatar:      item.Avatar,
				NickName:    item.NickName,
				Liked:       false,
			})
			i++
		}
	}

	// 按 idList 排序
	for i, v := range idList {
		for j, k := range listItem {
			if i >= len(listItem) || j >= len(idList) {
				continue
			}
			if v == k.Id {
				listItem[i], listItem[j] = listItem[j], listItem[i]
			}
		}
	}

	return listItem, nil
}

package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
)

func GetIdea(id, uid int) (*model.IdeaListItem, error) {
	item, err := model.GetIdeaById(id)
	if err != nil {
		return nil, err
	}

	if item.Privacy == 1 && item.UserId != uid {
		return nil, errno.ErrPermissionDenied
	}

	// 查询点赞
	likeList, err := model.GetIdeaLikeRecordForUser(uid, []int{item.Id})
	if err != nil {
		return nil, err
	}

	res := &model.IdeaListItem{
		Id:          item.Id,
		Content:     item.Content,
		ReleaseDate: item.ReleaseDate,
		LikesSum:    item.LikesSum,
		CommentSum:  item.CommentSum,
		UserId:      item.UserId,
		Avatar:      item.Avatar,
		NickName:    item.NickName,
	}

	if len(likeList) != 0 {
		res.Liked = true
	}

	return res, nil
}

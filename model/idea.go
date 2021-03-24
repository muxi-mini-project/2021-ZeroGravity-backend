package model

type IdeaModel struct {
	IdeaId      int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
	Content     string `json:"content" gorm:"column:content;" binding:"required"`
	ReleaseDate string `json:"releaseDate" gorm:"column:releaseDate;" binding:"required"`
	PublisherId int    `json:"publisher_id" gorm:"column:publisher_id;" binding:"required"`
	LikesSum    int    `json:"likes_sum" gorm:"column:likes_sum;" binding:"required"`
	CommentSum  int    `json:"comment_sum" gorm:"column:comment_sum;" binding:"required"`
	Privacy     int    `json:"privacy" gorm:"column:privacy;" binding:"required"` // 0->公有 1->私有
}

type IdeaInfo struct {
	Id          int    `json:"idea_id" gorm:"column:idea_id;" binding:"required"`
	Content     string `json:"content" gorm:"column:content;" binding:"required"`
	ReleaseDate string `json:"release_date" gorm:"column:release_date;" binding:"required"`
	LikesSum    int    `json:"likes_sum" gorm:"column:likes_sum;" binding:"required"`
	CommentSum  int    `json:"comment_sum" gorm:"column:comment_sum;" binding:"required"`
	Privacy     int    `json:"privacy" gorm:"column:privacy;" binding:"required"`
	UserId      int    `json:"publisher_id" gorm:"column:publisher_id;" binding:"required"`
	Avatar      string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	NickName    string `json:"nickname" gorm:"column:nickname;" binding:"required"`
}

type IdeaListItem struct {
	Id          int    `json:"idea_id"`
	Content     string `json:"content"`
	ReleaseDate string `json:"release_date"`
	LikesSum    int    `json:"likes_sum"`
	CommentSum  int    `json:"comment_sum"`
	UserId      int    `json:"publisher_id"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nickname"`
	Liked       bool   `json:"liked"` // 是否点赞
}

func (u *IdeaModel) TableName() string {
	return "tbl_idea"
}

func (u *IdeaModel) Create() error {
	return DB.Self.Create(u).Error
}

func DeleteIdea(id, uid int) error {
	u := &IdeaModel{}
	u.IdeaId = id
	d := DB.Self.Where("publisher_id = ?", uid).Delete(u)
	return d.Error
}

func GetIdeaById(id int) (*IdeaInfo, error) {
	u := &IdeaInfo{}

	query := DB.Self.Table("tbl_idea").
		Select("tbl_idea.*,tbl_user.nickname,tbl_user.avatar").
		Joins("left join tbl_user on tbl_user.id = tbl_idea.publisher_id").
		Where("tbl_idea.idea_id = ?", id)

	if err := query.Scan(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func GetIdeaList(uid, offset, limit, privicy, index, userId int) ([]*IdeaInfo, error) {
	item := make([]*IdeaInfo, 0)

	query := DB.Self.Table("tbl_idea").
		Select("tbl_idea.*,tbl_user.nickname,tbl_user.avatar").
		Joins("left join tbl_user on tbl_user.id = tbl_idea.publisher_id").
		Offset(offset).Limit(limit)

	// privicy 判断
	if privicy == 0 {
		// 获取他人想法
		query = query.Where("tbl_idea.privicy = ?", 0)

		// 判断是否选择用户
		if userId != 0 {
			query = query.Where("tbl_idea.publisher_id = ?", userId)
		}
	} else {
		// 获取用户自己的想法
		query = query.Where("tbl_idea.publisher_id = ?", uid)
	}

	// 判断排序顺序
	if index == 0 {
		query = query.Order("tbl_idea.idea_id desc")
	} else {
		query = query.Order("tbl_idea.likes_sum desc")
	}

	if err := query.Scan(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

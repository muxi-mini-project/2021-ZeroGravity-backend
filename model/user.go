package model

// --Model--

type UserModel struct {
	Id              int    `json:"id" gorm:"column:id;" binding:"required"`
	Account         string `json:"account" gorm:"column:account;" binding:"required"`
	AccountPassword string `json:"account_password" gorm:"column:account_password;" binding:"required"`
	NickName        string `json:"nickname" gorm:"column:nickname;" binding:"required"`
	Avatar          string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	Energy          int    `json:"energy" gorm:"column:energy;" binding:"required"`
}

func (u *UserModel) TableName() string {
	return "tbl_user"
}

func (u *UserModel) Create() error {
	return DB.Self.Create(u).Error
}

func GetUserByAccount(account string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Table("tbl_user").Where("account = ?", account).First(u)
	return u, d.Error
}

func GetUserByAccountAndPassword(account, accountPassword string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Table("tbl_user").Where("account = ? AND account_password = ?", account, accountPassword).First(u)
	return u, d.Error
}

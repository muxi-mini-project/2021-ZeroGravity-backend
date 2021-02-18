package model

// --Model--

type UserModel struct {
	Id              string `json:"id" gorm:"column:id;" binding:"required"`
	Account         string `json:"account" gorm:"column:account;" binding:"required"`
	AccountPassword string `json:"account_password" gorm:"column:account_password;" binding:"required"`
	NickName        string `json:"nickname" gorm:"column:nickname;" binding:"required`
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

func GetUserByAccountAndPassword(req *LoginRequest) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Table("tbl_user").Where("account = ? AND account_password = ?", req.Account, req.AccountPassword).First(u)
	return u, d.Error
}

// --Request&Response--

type CreateUserRequest struct {
	Account         string `json:"account"`
	AccountPassword string `json:"account_password"`
	NickName        string `json:"nickname"`
	Avatar          string `json:"avatar"`
}

type LoginRequest struct {
	Account         string `json:"account"`
	AccountPassword string `json:"account_password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

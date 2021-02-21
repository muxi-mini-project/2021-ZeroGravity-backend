package user

type CreateUserRequest struct {
	Account         string `json:"account"`
	AccountPassword string `json:"account_password"`
	NickName        string `json:"nickname"`
	Avatar          string `json:"avatar"`
}

type GetUserInfoResponse struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
	Energy   int    `json:"energy"`
}

type UpdateUserInfoRequest struct {
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
}

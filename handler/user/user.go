package user

// GetUserInfoResponse ... 单个用户
type GetUserInfoResponse struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
	Energy   int    `json:"energy"`
}

// UpdateUserInfoRequest ... 更改用户信息请求
type UpdateUserInfoRequest struct {
	Avatar   string `json:"avatar"`
	NickName string `json:"nickname"`
}

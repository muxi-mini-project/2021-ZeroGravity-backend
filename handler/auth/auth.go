package auth

// LoginRequest ... 登陆请求，密码需要 base64 加密
type LoginRequest struct {
	Account         string `json:"account"`
	AccountPassword string `json:"account_password"`
}

// LoginResponse ... 登陆响应，返回 token
type LoginResponse struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

// CreateUserRequest ... 注册请求,密码需要 base64 加密
type CreateUserRequest struct {
	Account         string `json:"account"`
	AccountPassword string `json:"account_password"`
	NickName        string `json:"nickname"`
	Avatar          string `json:"avatar"`
}

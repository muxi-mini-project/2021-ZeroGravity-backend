package auth

type LoginRequest struct {
	Account         string `json:"account"`
	AccountPassword string `json:"account_password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CreateUserRequest struct {
	Account         string `json:"account"`
	AccountPassword string `json:"account_password"`
	NickName        string `json:"nickname"`
	Avatar          string `json:"avatar"`
}

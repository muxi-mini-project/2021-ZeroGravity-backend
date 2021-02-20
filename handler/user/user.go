package user

import (
	"github.com/2021-ZeroGravity-backend/model"
)

type IdeaResponse struct {
	Count int                   `json:"count"`
	List  []*model.IdeaListItem `json:"list"`
}

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

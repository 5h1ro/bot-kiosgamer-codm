package users

import (
	"botkiosgamercodm/entity"
	"botkiosgamercodm/libs"
)

type Repository interface {
	CreateUser(game_id, session_key string) entity.UserResponse
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r repository) CreateUser(game_id, session_key string) entity.UserResponse {
	datadome, cookies := libs.GetDatadome()
	loginData := libs.Login(game_id, datadome, cookies)
	_ = loginData
	libs.LoginSSO(session_key, cookies)
	// username := res["username"].(string)
	// csrf := libs.GetCsrf(session_key)
	// libs.Redeem(session_key, username, csrf, code)
	// println(username, csrf)
	return entity.UserResponse{
		IsError: true,
	}
}

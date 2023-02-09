package users

import "botkiosgamercodm/entity"

type Service interface {
	CreateUserService(game_id, session_key string) entity.UserResponse
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUserService(game_id, session_key string) entity.UserResponse {
	res := s.repository.CreateUser(game_id, session_key)
	return res
}

package service

import (
	"Siberat/entity"
	"Siberat/repository"
)

type UserService interface {
	GetAllUser() ([]entity.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		repository: r,
	}
}

func(s *userService) GetAllUser() ([]entity.User, error){
	return s.repository.GetAllUser()
}
package service

import (
	"Siberat/dto"
	"Siberat/entity"
	errorhandler "Siberat/errorHandler"
	"Siberat/repository"
	"strconv"
)

type UserService interface {
	GetAllUser() ([]entity.User, error)
	UpdateUser(id int, req *dto.UpdateUserRequest) error
	GetUserByID(id int)(*entity.User, error)
	DeleteUser(id int) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		repository: r,
	}
}

func (s *userService) GetAllUser() ([]entity.User, error) {
	return s.repository.GetAllUser()
}


func (s *userService) GetUserByID(id int) (*entity.User, error){
	return s.repository.GetUserByID(id)
}

func (s *userService) UpdateUser(id int, req *dto.UpdateUserRequest) error {
	roleID, err := strconv.Atoi(req.IDRole)
	if err != nil {
		return &errorhandler.BadRequestError{Message: "invalid role id"}
	}

	user := entity.User{
		Name:        req.Name,
		UserName:    req.UserName,
		Email:       req.Email,
		NoHp:        req.NoHp,
		IDRole:      roleID,
		KodeWilayah: req.KodeWilayah,
	}
	return s.repository.UpdateUser(id, user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repository.DeleteUser(id)
}

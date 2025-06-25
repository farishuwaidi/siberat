package service

import (
	"Siberat/dto"
	"Siberat/entity"
	errorhandler "Siberat/errorHandler"
	"Siberat/helper"
	"Siberat/repository"
	"strconv"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	// cek email pada database
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "email alredy exist"}
	}

	passwrodHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	roleID, err := strconv.Atoi(req.Role)
	if err != nil {
		return &errorhandler.BadRequestError{Message: "Invalid role ID"}
	}

	user := entity.User {
		Name: req.Name,
		UserName: req.UserName,
		Email: req.Email,
		NoHp: req.NoHp,
		Password: passwrodHash,
		RoleID: roleID,
		KodeWilayah: req.KodeWilayah,
	}
	
	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Email not found"}
	}

	s.repository.PreLoadUserRole(user)

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong Password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID: user.ID,
		Name: user.Name,
		Token: token,
		Role: user.Role.Role,
	}
	return &data, nil
}
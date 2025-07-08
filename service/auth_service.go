package service

import (
	"Siberat/dto"
	"Siberat/entity"
	errorhandler "Siberat/errorHandler"
	"Siberat/helper"
	"Siberat/repository"
	"fmt"
	"strconv"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetPermissionData(id string) (*dto.GetPermissionResponse, error)
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
		return &errorhandler.BadRequestError{MessageText: "email alredy exist"}
	}

	passwrodHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{MessageText: err.Error()}
	}
	roleID, err := strconv.Atoi(req.Role)
	if err != nil {
		return &errorhandler.BadRequestError{MessageText: "Invalid role ID"}
	}

	user := entity.User {
		Name: req.Name,
		UserName: req.UserName,
		Email: req.Email,
		NoHp: req.NoHp,
		Password: passwrodHash,
		IDRole: roleID,
		KodeWilayah: req.KodeWilayah,
	}
	
	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{MessageText: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByUsername(req.UserName)
	if err != nil {
		return nil, &errorhandler.NotFoundError{MessageText: "User not found"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{MessageText: "Wrong Password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{MessageText: err.Error()}
	}

	data = dto.LoginResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		UserName: user.UserName,
		EmailVerifiedAt: nil,
		UserData: "",
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		KodeWilayah: user.KodeWilayah,
		KodeWilayahProses: user.KodeWilayahProses,
		NoHp: user.NoHp,
		NoWa: user.NoWa,
		UserNameSipandu: user.UserNameSipandu,
		UserPasswordSipandu: user.UserPasswordSipandu,
		IDRole: user.IDRole,
		Token: token,
		TokenExpiredAt: helper.GetTokenExpirationTime(),
		PhotoUrl: "https://ui-avatars.com/api/" ,
	}
	return &data, nil
}

func (s *authService) GetPermissionData(id string) (*dto.GetPermissionResponse, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("error user tidak ditemukan: %v", err)
	}

	role, err := s.repository.GetRoleByID(user.IDRole)
	if err != nil {
		return nil, fmt.Errorf("error role tidak ditemukan: %v", err)
	}

	toStr := func(v int) string {
		return fmt.Sprintf("%d", v)
	}

	return &dto.GetPermissionResponse{
		IdRole: role.IdRole,
		NmRole: role.NmRole,
		StsDaftar: toStr(role.StsDaftar),
		StsTambahPotensi: toStr(role.StsTambahPotensi),
		StsInfoPab: toStr(role.StsInfoPab),
		StsKirimUlgSkkp: toStr(role.StsKirimUlgSkkp),
		StsCetakNab: toStr(role.StsCetakNab),
		StsCetakUlgNab: toStr(role.StsCetakUlgNab),
		StsInfoSubjekPab: toStr(role.StsInfoSubjekPab),
		StsDaftarUlgPotBaru: toStr(role.StsDaftarUlgPotBaru),
		StsUserMngmt: toStr(role.StsUserMngmt),
		StsNjabMngmt: toStr(role.StsNjabMngmt),
		StsInfoNjab: toStr(role.StsInfoNjab),
		StsInfoUser: toStr(role.StsInfoUser),
		StsTable: toStr(role.StsTable),
		StsTableMngmt: toStr(role.StsTableMngmt),
		StsVerif: toStr(role.StsVerif),
	}, nil
}
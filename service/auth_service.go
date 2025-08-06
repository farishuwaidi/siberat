package service

import (
	"Siberat/dto"
	errorhandler "Siberat/errorHandler"
	"Siberat/helper"
	"Siberat/repository"
	"fmt"
)

type AuthService interface {
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetPermissionData(id string) (*dto.GetPermissionResponse, error)
	UpdatePassword(req *dto.UpdatePassword) (*dto.UpdatePasswordResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
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
		KodeWilayah: user.KdWil,
		KodeWilayahProses: user.KdWilKerja,
		NoHp: user.NoHp,
		NoWa: user.NoWa,
		UserNameSipandu: user.UsernameSipandu,
		UserPasswordSipandu: user.UserpasswordSipandu,
		IDRole: user.IdRole,
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

	role, err := s.repository.GetRoleByID(user.IdRole)
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

func (s *authService) UpdatePassword(req *dto.UpdatePassword) (*dto.UpdatePasswordResponse, error) {
	user, err := s.repository.GetUserByID(fmt.Sprintf("%d", req.ID))
	if err != nil {
		return nil, &errorhandler.NotFoundError{ MessageText: "user tidak ditemukan"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.BadRequestError{MessageText: "password lama tidak sesuai"}
	}

	if req.PasswordBaru != req.PasswordConfirm {
		return nil, &errorhandler.BadRequestError{ MessageText: "konfirmasi password tidak sesuai"}
	}

	hashPassword := helper.HashPassword(req.PasswordBaru)

	if err := s.repository.UpdatePassword(user.ID, hashPassword); err != nil {
		return nil, &errorhandler.InternalServerError{MessageText: "gagal update password"}
	}

	return &dto.UpdatePasswordResponse{Message: "password berhasil diubah"}, nil
}
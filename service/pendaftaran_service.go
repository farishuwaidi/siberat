package service

import (
	"Siberat/dto"
	"Siberat/repository"
)

type PendaftaranService interface {
	CekTransaksi(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.CekTransaksiResponse, error)
}

type pendaftaranService struct {
	repository repository.PendaftaranRepository
}

func NewPendaftaranService(r repository.PendaftaranRepository) *pendaftaranService {
	return &pendaftaranService{
		repository: r,
	}
}

func (s *pendaftaranService) CekTransaksi(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.CekTransaksiResponse, error) {
	return s.repository.CekTransaksi(NoAb1,NoAb2,NoAb3)
}
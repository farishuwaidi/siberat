package service

import (
	"Siberat/dto"
	"Siberat/repository"
)

type ReferenceService interface{
	GetKodeRole() ([]dto.ResponseKodeRole, error)
	GetKodeWil() ([]dto.ResponseKodeWil,error)
}

type referenceService struct {
	repository repository.ReferenceRepository
}

func NewReferenceService(r repository.ReferenceRepository) *referenceService {
	return &referenceService{
		repository: r,
	}
}

func (s *referenceService) GetKodeRole() ([]dto.ResponseKodeRole, error) {
	return s.repository.GetKodeRole()
}

func (s *referenceService) GetKodeWil() ([]dto.ResponseKodeWil, error) {
	return s.repository.GetKodeWil()
}
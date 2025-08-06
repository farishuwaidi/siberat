package service

import (
	"Siberat/dto"
	"Siberat/repository"
)

type SubjekService interface{
	GetSubjekPajak(field, value string) (*dto.ResponseGetSubjek, error)
	GetKepemilikanAb(id string) ([]dto.GetKepemilikanAbResponse, error)
}

type subjekService struct {
	repository repository.SubjekRepository
}

func NewSubjekService(r repository.SubjekRepository) *subjekService {
	return &subjekService{r}
}

func (s *subjekService) GetSubjekPajak(field, value string) (*dto.ResponseGetSubjek, error) {
	return s.repository.GetSubjekPajak(field, value)
}

func (s *subjekService) GetKepemilikanAb(id string) ([]dto.GetKepemilikanAbResponse, error) {
	return s.repository.GetKepemilikanAb(id)
} 
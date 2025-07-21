package service

import (
	"Siberat/dto"
	"Siberat/repository"
)

type ObjekService interface{
	GetObjekByNab(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.ObjekPabResponse, error)
}

type objekService struct {
	repository repository.ObjekRepository
}

func NewObjeKService(r repository.ObjekRepository) *objekService {
	return &objekService{
		repository: r,
	}
}

func (s *objekService) GetObjekByNab(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.ObjekPabResponse, error) {
	return s.repository.GetObjekByNab(NoAb1, NoAb2, NoAb3)
}
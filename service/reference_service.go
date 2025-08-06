package service

import (
	"Siberat/dto"
	"Siberat/repository"
)

type ReferenceService interface{
	GetKodeRole() ([]dto.ResponseKodeRole, error)
	GetKodeWil() ([]dto.ResponseKodeWil,error)
	GetKdJenisAB() ([]dto.ResponseKdJenisAB, error)
	GetKdMerek(kdJenisAb string) ([]dto.ResponseKdMerek, error)
	GetKdModel(KdMerekAb string) ([]dto.ResponseKdModel, error)
	GetNilaiJual(kdMerek string) ([]dto.ResponseNilaiJual, error)
	GetJenisBbm() ([]dto.ResponseJenisBBM, error)
	GetJenisKepemilikan() ([]dto.ResponseJenisKepemilikan, error)
	GetKdMohon() ([]map[string]interface{}, error)
	
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

func (s *referenceService) GetKdJenisAB() ([]dto.ResponseKdJenisAB, error) {
	return s.repository.GetKdJenisAB()
}

func (s *referenceService) GetKdMerek(kdJenisAb string) ([]dto.ResponseKdMerek, error){
	return s.repository.GetKdMerek(kdJenisAb)
}

func (s *referenceService) GetKdModel(KdMerekAb string) ([]dto.ResponseKdModel, error) {
	return s.repository.GetKdModel(KdMerekAb)
}

func (s *referenceService) GetNilaiJual(kdMerek string) ([]dto.ResponseNilaiJual, error) {
	return s.repository.GetNilaiJual(kdMerek)
}

func (s *referenceService) GetJenisBbm() ([]dto.ResponseJenisBBM, error) {
	return s.repository.GetJenisBbm()
}

func (s *referenceService) GetJenisKepemilikan() ([]dto.ResponseJenisKepemilikan, error) {
	return s.repository.GetJenisKepemilikan()
}

func (s *referenceService) GetKdMohon() ([]map[string]interface{}, error) {
	return s.repository.GetKdMohon()
}
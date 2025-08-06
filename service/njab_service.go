package service

import (
	"Siberat/dto"
	"Siberat/repository"
	"errors"
	"fmt"
	"time"
)

type NjabService interface{
	UpdateMerek(req dto.UpdateNjabRequest) (map[string]interface{}, error)
	GetMerekDetails(kdMerek *dto.GetNjabDetailsRequest) (map[string]interface{}, error)
	UpdateNilaiJual(req *dto.UpdateNjabRequest) (map[string]interface{}, error)
	GetDetailNjab(kdMerek string, thBuatan string) (map[string]interface{}, error)
}

type njabService struct {
	repository repository.NjabRepository
	repouser repository.UserRepository
}

func NewNjabService(r repository.NjabRepository, ru repository.UserRepository) *njabService {
	return &njabService{
		repository: r,
		repouser:   ru,
	}
}

func (s *njabService) UpdateMerek(req dto.UpdateNjabRequest) (map[string]interface{}, error) {
	switch req.Method {
	case "A":
		merek := req.DataSent
		user := req.IdUser
		userData, err := s.repouser.GetUserDetail(user)
		if err !=nil {
			return nil, fmt.Errorf("user tidak ditemukan: %s", user)
		}

		newMerek := map[string]interface{}{
			"kd_merek_ab": merek["kd_merek_ab"],
			"nm_merek_ab": merek["nm_merek_ab"],
			"nm_model_ab": merek["nm_model_ab"],
			"edit_by": userData.SurName,
		}

		if err := s.repository.AddMerek(newMerek); err != nil {
			return nil, err
		}
		reqDetails := &dto.GetNjabDetailsRequest{
			KdMerekAb: fmt.Sprintf("%v", merek["kd_merek_ab"]),
		}
		return s.repository.GetMerekDetails(reqDetails)
	case "U":
		kdMerekAb := req.DataSent["kd_merek_ab"]
		user := req.IdUser
		userData, err := s.repouser.GetUserDetail(user)
		if err != nil {
			return nil, fmt.Errorf("user tidak ditemukan: %s", user)
		}
		dataEdit := map[string]interface{}{
			"nm_merek_ab": req.DataSent["nm_merek_ab"],
			"nm_model_ab": req.DataSent["nm_model_ab"],
			"edit_by": userData.SurName,
			"updated_at": time.Now(),
		}
		if err := s.repository.UpdateMerek(kdMerekAb, dataEdit); err != nil {
			return nil, err
		}

		kdMerekStr, ok := kdMerekAb.(string)
		if !ok {
			return nil, errors.New("kdMerekAb is not string")
		}
		reqDetails := &dto.GetNjabDetailsRequest{
			KdMerekAb: kdMerekStr,
		}
		return s.repository.GetMerekDetails(reqDetails)
	case "D":
		kdMerekAb := req.DataSent["kd_merek_ab"]
		if err := s.repository.DeleteMerek(kdMerekAb); err != nil {
			return nil, err
		}
		return req.DataSent, nil
	}
	return nil, errors.New("invalid method")
}

func (s *njabService) GetMerekDetails(req *dto.GetNjabDetailsRequest) (map[string]interface{}, error) {
	return s.repository.GetMerekDetails(req)
}

func (s *njabService) UpdateNilaiJual(req *dto.UpdateNjabRequest) (map[string]interface{}, error) {
	user := req.IdUser
	userData, err := s.repouser.GetUserDetail(user)
	if err != nil {
		return nil, fmt.Errorf("user tidak ditemukan: %s", user)
	}
	nlJual := req.DataSent
	method := req.Method

	switch method {
	case "A":
		required := []string{"kd_merek_ab", "th_buatan", "nilai_jual", "bobot", "pab_dasar"}
		for _, field := range required {
			if _, exist := nlJual[field]; !exist {
				return nil, fmt.Errorf("field %s is required", field)
			}
		}
		found, err := s.repository.FindMerek(nlJual["kd_merek_ab"].(string))
		if err != nil || !found {
			return nil, fmt.Errorf("error checking merek existence: %v", err)
		}

		nilaiJual := map[string]interface{}{
			"kd_merek_ab": nlJual["kd_merek_ab"],
			"th_buatan":   nlJual["th_buatan"],
			"nilai_jual": nlJual["nilai_jual"],
			"bobot":       nlJual["bobot"],
			"pab_dasar":   nlJual["pab_dasar"],
			"edit_by":     userData.SurName,
			"updated_at": time.Now(),
		}

		if err := s.repository.AddNilaiJual(nilaiJual); err != nil {
			return nil, fmt.Errorf("error adding nilai jual: %v", err)
		}
		
		reqDetails := &dto.GetNilaiJualRequest{
			KdMerekAb: nlJual["kd_merek_ab"].(string),
			ThBuatan:  nlJual["th_buatan"].(string),
		}
		return s.repository.GetDetailNjab(reqDetails.KdMerekAb, reqDetails.ThBuatan)
	case "U":
		kdMerekAb := nlJual["kd_merek_ab"].(string)
		thBuatan := nlJual["th_buatan"].(string)

		dataEdit, ok := nlJual["data_edit"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("data_edit is required")
		}
		dataEdit["edit_by"] = userData.SurName
		dataEdit["updated_at"] = time.Now()
		err := s.repository.UpdateNilaiJual(kdMerekAb, thBuatan, dataEdit)
		if err != nil {
			return nil, fmt.Errorf("error updating nilai jual: %v", err)
		}
		return s.repository.GetDetailNjab(kdMerekAb, thBuatan)
	case "D":
		kdMerek := nlJual["kd_merek_ab"].(string)
		thBuatan := nlJual["th_buatan"].(string)

		err := s.repository.DeleteNilaiJual(kdMerek, thBuatan)
		if err != nil {
			return nil, fmt.Errorf("error deleting nilai jual: %v", err)
		}
		return req.DataSent, nil
	default:
		return nil, errors.New("invalid method")
	}
}

func (s *njabService) GetDetailNjab(kdMerek string, thBuatan string) (map[string]interface{}, error) {
	return s.repository.GetDetailNjab(kdMerek, thBuatan)
}
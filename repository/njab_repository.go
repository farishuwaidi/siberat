package repository

import (
	"Siberat/dto"

	"gorm.io/gorm"
)

type NjabRepository interface {
	GetMerekDetails(kdMerek *dto.GetNjabDetailsRequest) (map[string]interface{}, error)
	AddMerek(merek map[string]interface{}) error
	UpdateMerek(kdMerek any, merek map[string]interface{}) error
	DeleteMerek(kdMerek any) error
	AddNilaiJual(data map[string]interface{}) error
	UpdateNilaiJual(kdMerek string, thBuatan string, data map[string]interface{}) error
	DeleteNilaiJual(kdMerek string, thBuatan string) error
	FindMerek(kdMerek string) (bool, error)
	GetDetailNjab(kdMerek string, thBuatan string) (map[string]interface{}, error)
}

type njabRepository struct {
	db *gorm.DB
}

func NewNjabRepository(db *gorm.DB) *njabRepository {
	return &njabRepository{db: db}
}

func (r *njabRepository) GetMerekDetails(kdMerek *dto.GetNjabDetailsRequest) (map[string]interface{}, error) {
	var merekDetails map[string]interface{}

	err := r.db.Table("t_merekab").
		Where("kd_merek_ab = ?", kdMerek.KdMerekAb).
		Scan(&merekDetails).Error
	if err != nil {
		return nil, err
	}
	return merekDetails, nil
}
func (r *njabRepository) AddMerek(merek map[string]interface{}) error {
	return r.db.Table("t_merekab").Create(merek).Error
}

func (r *njabRepository) UpdateMerek(kdMerek any, merek map[string]interface{}) error {
	return r.db.Table("t_merekab").Where("kd_merek_ab = ?", kdMerek).Updates(merek).Error
}

func (r *njabRepository) DeleteMerek(kdMerek any) error {
	return r.db.Table("t_merekab").Where("kd_merek_ab = ?", kdMerek).Delete(nil).Error
}

func (r *njabRepository) AddNilaiJual(data map[string]interface{}) error {
	return r.db.Table("t_nljualab").Create(data).Error
}

func (r *njabRepository) UpdateNilaiJual(kdMerek string, thBuatan string, data map[string]interface{}) error {
	return r.db.Table("t_nljualab").
		Where("kd_merek_ab = ? AND th_buatan = ?", kdMerek, thBuatan).
		Updates(data).Error
}

func (r *njabRepository) DeleteNilaiJual(kdMerek string, thBuatan string) error {
	return r.db.Table("t_nljualab").
		Where("kd_merek_ab = ? AND th_buatan = ?", kdMerek, thBuatan).
		Delete(nil).Error
}

func (r *njabRepository) GetDetailNjab(kdMerek string, thBuatan string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := r.db.Table("t_nljualab").
		Where("kd_merek_ab = ? AND th_buatan = ?", kdMerek, thBuatan).
		Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *njabRepository) FindMerek(kdMerek string) (bool, error) {
	var count int64
	err := r.db.Table("t_merekab").
		Where("kd_merek_ab = ?", kdMerek).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
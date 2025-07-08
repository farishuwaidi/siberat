package repository

import (
	"Siberat/dto"

	"gorm.io/gorm"
)

type ReferenceRepository interface{
	GetKodeRole() ([]dto.ResponseKodeRole, error)
	GetKodeWil() ([]dto.ResponseKodeWil, error)
	
}

type referenceRepository struct {
	db *gorm.DB
}

func NewReferenceRepository(db *gorm.DB) *referenceRepository {
	return &referenceRepository{db: db}
}
func (r *referenceRepository) GetKodeRole() ([]dto.ResponseKodeRole, error) {
	var roles []dto.ResponseKodeRole
	err := r.db.
		Table("t_role").
		Select("id_role, nm_role").
		Scan(&roles).Error
	// err := r.db.Raw("SELECT CAST(id_role AS int) AS id_role, nm_role FROM t_role").Scan(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func  (r *referenceRepository) GetKodeWil() ([]dto.ResponseKodeWil, error) {
	var wilayah []dto.ResponseKodeWil
	err := r.db.
		Table("t_wiluppd").
		Select("kd_wil, nm_wil, al_wil, kab_kota, provinsi").
		Scan(&wilayah).Error
	if err != nil {
		return nil, err
	}
	return wilayah, nil
}

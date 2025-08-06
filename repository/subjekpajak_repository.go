package repository

import (
	"Siberat/dto"
	"fmt"

	"gorm.io/gorm"
)

type SubjekRepository interface{
	GetSubjekPajak(field string, value string) (*dto.ResponseGetSubjek, error)
	GetKepemilikanAb(id string) ([]dto.GetKepemilikanAbResponse, error)
}

type subjekRepository struct {
	db *gorm.DB
}

func NewSubjekRepository(db *gorm.DB) *subjekRepository {
	return &subjekRepository{db}
}

func (r *subjekRepository) GetSubjekPajak(field string, value string) (*dto.ResponseGetSubjek, error) {
	var result dto.ResponseGetSubjek

	// validate whitelist field
	allowedField := map[string]bool{
		"npwp":true, "nib": true, "no_ktp":true, "npwpd":true,
	}

	if !allowedField[field] {
		return nil, fmt.Errorf("invalid search field")
	}

	query := fmt.Sprintf("%s = ?", field)
	err := r.db.
		Table("t_subjek_pajak").
		Where(query, value).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	if result.IDSubjekPajak == "" {
		return nil, nil
	}

	return &result, nil
}

func (r *subjekRepository) GetKepemilikanAb(id string) ([]dto.GetKepemilikanAbResponse, error) {
	var listAb []dto.GetKepemilikanAbResponse

	err := r.db.
		Table("t_objek_pajak").
		Joins("LEFT JOIN t_merekab ON t_objek_pajak.kd_merek_ab = t_merekab.kd_merek_ab ").
		Joins("LEFT JOIN t_kdjenisab ON t_objek_pajak.kd_jenis_ab = t_kdjenisab.kd_jenis_ab").
		Joins("LEFT JOIN t_wiluppd ON t_objek_pajak.kd_wil = t_wiluppd.kd_wil").
		Select(`
			t_objek_pajak.no_ab1 AS no_ab1,
			t_objek_pajak.no_ab2 AS no_ab2,
			t_objek_pajak.no_ab3 AS no_ab3,
			t_objek_pajak.kd_wil AS kd_Wil,
			t_wiluppd.nm_wil AS nm_wil,
			t_merekab.nm_merek_ab AS nm_merek_ab,
			t_merekab.nm_model_ab AS nm_model_ab,
			t_objek_pajak.th_buatan AS th_buatan,
			t_kdjenisab.nm_jenis_ab AS nm_jenis_ab,
			TO_CHAR(t_objek_pajak.tg_akhir_pajak, 'YYYY-MM-DD') AS tg_akhir_pajak
		`).
		Where("id_subjek_pajak = ?", id).
		Scan(&listAb).Error;

	if err != nil {
		return nil, err
	}

	return listAb, err
		
}
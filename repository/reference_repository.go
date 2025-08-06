package repository

import (
	"Siberat/dto"

	"gorm.io/gorm"
)

type ReferenceRepository interface{
	GetKodeRole() ([]dto.ResponseKodeRole, error)
	GetKodeWil() ([]dto.ResponseKodeWil, error)
	GetKdJenisAB() ([]dto.ResponseKdJenisAB, error)
	GetKdMerek(kdJenisAb string) ([]dto.ResponseKdMerek, error)
	GetKdModel(KdMerekAb string) ([]dto.ResponseKdModel, error)
	GetNilaiJual(kdMerek string) ([]dto.ResponseNilaiJual, error)
	GetJenisBbm() ([]dto.ResponseJenisBBM, error)
	GetJenisKepemilikan() ([]dto.ResponseJenisKepemilikan, error)
	GetKdMohon() ([]map[string]interface{}, error)
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

func (r *referenceRepository) GetKdJenisAB() ([]dto.ResponseKdJenisAB, error) {
	var jenisAB []dto.ResponseKdJenisAB
	err := r.db.
		Table("t_kdjenisab").
		Select("kd_jenis_ab, nm_jenis_ab").
		Scan(&jenisAB).Error
	if err != nil {
		return nil, err
	}
	return jenisAB, nil
}

func (r *referenceRepository) GetKdMerek(kdJenisAb string) ([]dto.ResponseKdMerek, error) {
	var listMerek []dto.ResponseKdMerek
	// prefix := KdJenisAB + "%"
	err := r.db.
		Table("t_merekab").
		Select("kd_merek_ab, nm_merek_ab").
		Where("LEFT(kd_merek_ab, 3) = ?",kdJenisAb).
		Group("kd_merek_ab, nm_merek_ab").
		Order("kd_merek_ab").
		Scan(&listMerek).Error;
	if err != nil {
		return nil, err
	}
	return listMerek, nil
}

func (r *referenceRepository) GetKdModel(kdMerekAb string) ([]dto.ResponseKdModel, error) {
	var  listModel []dto.ResponseKdModel
	err := r.db.
		Table("t_merekab").
		Select("kd_merek_ab, nm_merek_ab, nm_model_ab").
		Where("kd_merek_ab = ?", kdMerekAb).
		Group("kd_merek_ab, nm_merek_ab, nm_model_ab").
		Order("kd_merek_ab").
		Scan(&listModel).Error;
	if err != nil {
		return nil, err
	}

	return listModel, nil
}

func (r *referenceRepository) GetNilaiJual(kdMerek string) ([]dto.ResponseNilaiJual, error) {
	var listNilaiJual []dto.ResponseNilaiJual
	err := r.db.
		Table("t_nljualab").
		Select("kd_merek_ab, th_buatan, nilai_jual, bobot, pab_dasar, edit_by, updated_at").
		Where("kd_merek_ab = ?", kdMerek).
		Scan(&listNilaiJual).Error;
	
	if err != nil {
		return nil, err
	}
	
	return listNilaiJual, nil
}

func (r *referenceRepository) GetJenisBbm() ([]dto.ResponseJenisBBM, error) {
	var bbm []dto.ResponseJenisBBM
	err := r.db.
		Table("t_bbm").
		Select("id, nm_bbm").
		Scan(&bbm).Error;
	
	if err != nil {
		return nil, err
	}

	return bbm, nil
}

func (r *referenceRepository) GetJenisKepemilikan() ([]dto.ResponseJenisKepemilikan, error) {
	var kepemilikan []dto.ResponseJenisKepemilikan
	err := r.db.
		Table("t_kdkepemilikan").
		Select("kd_kepemilikan, nm_kepemilikan").
		Scan(&kepemilikan).Error;
	
	if err != nil {
		return nil, err
	}

	return kepemilikan, err
}

func (r *referenceRepository) GetKdMohon() ([]map[string]interface{}, error) {
	var listKdMohon []map[string]interface{}

	err := r.db.
		Table("t_kdmohon").
		Select(`*`).
		Scan(&listKdMohon).Error;

	if err != nil {
		return nil, err
	}

	return listKdMohon, err
}
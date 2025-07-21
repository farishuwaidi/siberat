package repository

import (
	"Siberat/dto"

	"gorm.io/gorm"
)

type ObjekRepository interface{
	GetObjekByNab(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.ObjekPabResponse, error)
}

type objekRepository struct {
	db *gorm.DB
}

func NewObjekRepository(db *gorm.DB) *objekRepository {
	return &objekRepository{db}
}

func (r *objekRepository) GetObjekByNab(NoAb1 string, NoAb2 string, NoAb3 string) (*dto.ObjekPabResponse, error) {
	var objek *dto.ObjekPabResponse
	err := r.db.
		Table("t_objek_pajak").
		Select(`t_objek_pajak.id_objek_pajak,
			t_objek_pajak.no_ab1,
			t_objek_pajak.no_ab2,
			t_objek_pajak.no_ab3,
			t_objek_pajak.kd_jenis_ab,
			t_subjek_pajak.no_ktp,
			t_subjek_pajak.npwp,
			t_subjek_pajak.nib,
			t_subjek_pajak.nm_pemilik,
			t_subjek_pajak.al_pemilik,
			t_subjek_pajak.kd_pos,
			kelurahan_subjek.kd_kelurahan AS kd_kelurahan_subjek_pajak,
			kelurahan_subjek.nm_kelurahan AS nm_kelurahan_subjek_pajak,
			kelurahan_subjek.kd_kecamatan AS kd_kecamatan_subjek_pajak,
			kelurahan_subjek.nm_kecamatan AS nm_kecamatan_subjek_pajak,
			wiluppd_subjek.kd_wil AS kd_wil_subjek_pajak,
			wiluppd_subjek.nm_wil AS nm_wil_subjek_pajak,
			wiluppd_subjek.kab_kota AS kab_kota_subjek_pajak,
			wiluppd_subjek.provinsi AS provinsi_subjek_pajak,
			wiluppd_objek.kd_wil AS kd_wil_objek_pajak,
			wiluppd_objek.nm_wil AS nm_wil_objek_pajak,
			wiluppd_objek.kab_kota AS kab_kota_objek_pajak,
			wiluppd_objek.provinsi AS provinsi_objek_pajak,
			t_kdjenisab.kd_jenis_ab,
			t_kdjenisab.nm_jenis_ab,
			t_merekab.kd_merek_ab,
			t_merekab.nm_merek_ab,
			t_objek_pajak.th_buatan,
			t_objek_pajak.no_rangka,
			t_objek_pajak.no_mesin,
			t_objek_pajak.jumlah_cc,
			t_objek_pajak.tg_akhir_pajak,
			t_file.id_ref AS foto_ab
		`).
		Joins("LEFT JOIN t_subjek_pajak ON t_objek_pajak.id_subjek_pajak = t_subjek_pajak.id_subjek_pajak").
		Joins("LEFT JOIN t_wiluppd AS wiluppd_objek ON t_objek_pajak.kd_wil = wiluppd_objek.kd_wil").
		Joins("LEFT JOIN t_wiluppd AS wiluppd_subjek ON t_subjek_pajak.kd_wil_subjek_pajak = wiluppd_subjek.kd_wil").
		Joins("LEFT JOIN t_kdkelurahan AS kelurahan_subjek ON t_subjek_pajak.kd_kelurahan = kelurahan_subjek.kd_kelurahan").
		Joins("LEFT JOIN t_kdjenisab ON t_objek_pajak.kd_jenis_ab = t_kdjenisab.kd_jenis_ab").
		Joins("LEFT JOIN t_merekab  ON t_objek_pajak.kd_merek_ab = t_merekab.kd_merek_ab").
		Joins("LEFT JOIN t_file ON t_objek_pajak.id_objek_pajak = t_file.id_ref").
		Where("t_objek_pajak.no_ab1 = ?", NoAb1).
		Where("t_objek_pajak.no_ab2 = ?", NoAb2).
		Where("t_objek_pajak.no_ab3 = ?", NoAb3).
		Scan(&objek).Error
	
	if err != nil {
		return nil, err
	}
	
	return objek, nil
}
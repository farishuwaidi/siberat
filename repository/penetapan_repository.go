package repository

import (
	"Siberat/dto"
	"Siberat/helper"
	"fmt"

	"gorm.io/gorm"
)

type PenetapanRepository interface{
	GetPenetapan(IdTrnab string)(*dto.GetTrnabResponse, error)
}

type penetapanRepository struct {
	db *gorm.DB
}

func NewPenetapanRepository(db *gorm.DB) *penetapanRepository{
	return &penetapanRepository{db}
}

func (r *penetapanRepository) GetPenetapan(IdTrnab string) (*dto.GetTrnabResponse, error) {
	var result *dto.GetTrnabResponse
	err := r.db.
		Table("t_trnab").
		Select(`
			id_trnab, kd_mohon1, kd_mohon2, kd_mohon3, kd_mohon4, kd_mohon5, kd_mohon6, kd_status, kd_wil_proses, tg_daftar, jam_daftar, id_user_daftar, nm_user_daftar, tg_tetap, jam_tetap, id_user_tetap, nm_user_tetap, tg_bayar, jam_bayar, id_user_bayar, nm_user_bayar, kd_channel_bayar, bea_bbnab_pok, bea_bbnab_den, bea_pab_pok, bea_pab_pok1, bea_pab_pok2, bea_pab_pok3, bea_pab_pok4, bea_pab_pok5, bea_pab_den, bea_pab_den1, bea_pab_den2, bea_pab_den3, bea_pab_den4, bea_pab_den5, bea_bbnab_pok_d, bea_bbnab_den_d, bea_pab_pok_d, bea_pab_pok1_d, bea_pab_pok2_d, bea_pab_pok3_d, bea_pab_pok4_d, bea_pab_pok5_d, bea_pab_den_d, bea_pab_den1_d, bea_pab_den2_d, bea_pab_den3_d, bea_pab_den4_d, bea_pab_den5_d, tg_akhir_pajak, tg_akhir_pajak_baru, tg_akhir_pajak_lama, tg_faktur, tg_fiskal, tg_gntmesin, tg_kwitansi, tg_uben, tg_blockir, no_ab1, no_ab2, no_ab3, id_objek_pajak, kd_merek_ab, nm_merek_ab, nm_model_ab, kd_jenis_ab, th_buatan, no_rangka, no_mesin, warna_ab, id_bbm, jumlah_cc, kd_wil, kd_kepemilikan, kd_blockir, tg_kepemilikan, jam_kepemilikan, id_subjek_pajak, no_ktp, npwp, nib, nm_pemilik, al_pemilik, kd_wil_subjek_pajak, kd_kelurahan, kd_kecamatan, rt, rw, id_subjek_pajak_type, email, no_wa, no_hp, kd_pos, is_able_bayar, nilai_jual, bobot, bayar_kdpn, ket1, ket2, ket3, selang_hari, selang_bulan, selang_tahun, updated_at, tg_koreksi, jam_koreksi, id_user_koreksi, nm_user_koreksi, no_ab1_lm, no_ab2_lm, no_ab3_lm, kd_merek_ab_lm, nm_merek_ab_lm, nm_model_ab_lm, th_buatan_lm, nilai_jual_lm, bobot_lm, kd_jenis_ab_lm, warna_ab_lm, jumlah_cc_lm, kd_kepemilikan_lm, kd_blockir_lm, is_copied, kd_cetak_tnab, tg_cetak_tnab, jam_cetak_tnab, id_user_cetak_tnab, nm_user_cetak_tnab, kode_bayar, tagihan_pokok_tot, tagihan_denda_tot, kode_bayar_expired, tagihan_no_ab, tagihan_tot, unit_kode, is_pab_editable
		`).
		Where("id_trnab = ?", IdTrnab).
		Scan(&result).Error
	
	if IdTrnab == "" {
		return nil, fmt.Errorf("invalid param")
	}
	if result == nil {
		return nil, fmt.Errorf("data not found")
	}
	if err != nil {
		return nil, err
	}

	result.TgDaftar = helper.FormatDate(result.TgDaftar)
	result.TgTetap = helper.FormatDate(result.TgTetap)
	result.TgBayar = helper.FormatDate(result.TgBayar)
	result.TgAkhirPajak = helper.FormatDate(result.TgAkhirPajak)
	result.TgAkhirPajakBaru = helper.FormatDate(result.TgAkhirPajakBaru)
	result.TgAkhirPajakLama = helper.FormatDate(result.TgAkhirPajakLama)
	result.TgKepemilikan = helper.FormatDate(result.TgKepemilikan)
	result.UpdatedAt = helper.FormatDate(result.UpdatedAt)
	return result, nil
}
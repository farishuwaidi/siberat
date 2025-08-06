package dto

type ObjekPabRequest struct {
	NoAb1 string `form:"no_ab1" json:"no_ab1" binding:"required"`
	NoAb2 string `form:"no_ab2" json:"no_ab2" binding:"required"`
	NoAb3 string `form:"no_ab3" json:"no_ab3" binding:"required"`
}

type ObjekPabResponse struct {
	IdObjekPajak           string `json:"id_objek_pajak"`
	NoAb1                  string `json:"no_ab1"`
	NoAb2                  string `json:"no_ab2"`
	NoAb3                  string `json:"no_ab3"`
	KdJenisAb              string `json:"kd_jenis_ab"`
	NoKtp                  string `json:"kd_ktp"`
	Npwp                   string `json:"npwp"`
	Nib                    string `json:"nib"`
	NmPemilik              string `json:"nm_pemilik"`
	AlPemilik              string `json:"al_pemilik"`
	KdPos                  string `json:"kd_pos"`
	KdKelurahanSubjekPajak string `json:"kd_kelurahan_subjek_pajak"`
	NmKelurahanSubjekPajak string `json:"nm_kelurahan_subjek_pajak"`
	KdKecamatanSubjekPajak string `json:"kd_kecamatan_subjek_pajak"`
	NmKecamatanSubjekPajak string `json:"nm_kecamatan_subjek_pajak"`
	KdWilSubjekPajak       string `json:"kd_wil_subjek_pajak"`
	NmWilSubjekPajak       string `json:"nm_wil_subjek_pajak"`
	KabKotaSubjekPajak     string `json:"kab_kota_subjek_pajak"`
	ProvinsiSubjekPajak    string `json:"provinsi_subjek_pajak"`
	KdWilObjekPajak        string `json:"kd_wil_objek_pajak"`
	NmWilObbjekPajak       string `json:"nm_wil_objek_pajak"`
	KabKotaObjekPajak      string `json:"kab_kota_Objek_pajak"`
	ProvinsiObjekPajak     string `json:"provinsi_objek_pajak"`
	NmJenisAb              string `json:"nm_jenis_ab"`
	KdMerekAb              string `json:"kd_merek_ab"`
	NmMerekAb              string `json:"nm_merek_ab"`
	ThBuatan               string `json:"th_buatan"`
	NoRangka               string `json:"no_rangka"`
	NoMesin                string `json:"no_mesin"`
	JumlahCc               string `json:"jumlah_cc"`
	TgAkhirPajak           string `json:"tg_akhir_pajak"`
	FotoAb                 string `json:"foto_ab"`
}

type GetSubjekObjekPajak map[string]interface{}

type GetInfoSubjekObjek struct {
	IDObjekPajak           string  `json:"id_objek_pajak"`
	NoAB1                  string  `json:"no_ab1"`
	NoAB2                  string  `json:"no_ab2"`
	NoAB3                  string  `json:"no_ab3"`
	KdStatus               string  `json:"kd_status"`
	KdMerekAB              string  `json:"kd_merek_ab"`
	ThBuatan               string  `json:"th_buatan" gorm:"column:th_buatan"`
	KdMohon1               string  `json:"kd_mohon1"`
	KdMohon2               string  `json:"kd_mohon2"`
	KdMohon3               string  `json:"kd_mohon3"`
	KdMohon4               string  `json:"kd_mohon4"`
	KdMohon5               string  `json:"kd_mohon5"`
	KdMohon6               string  `json:"kd_mohon6"`
	NmMerekAB              string  `json:"nm_merek_ab"`
	NmModelAB              string  `json:"nm_model_ab"`
	NoRangka               string  `json:"no_rangka"`
	NoMesin                string  `json:"no_mesin"`
	WarnaAB                string  `json:"warna_ab"`
	IDBBM                  string  `json:"id_bbm"`
	JumlahCC               string  `json:"jumlah_cc"`
	TgAkhirPajak           string  `json:"tg_akhir_pajak" gorm:"column:tg_akhir_pajak"`
	IDSubjekPajak          string  `json:"id_subjek_pajak"`
	KdWil                  string  `json:"kd_wil"`
	KdWilProses            string  `json:"kd_wil_proses"`
	KdJenisAB              string  `json:"kd_jenis_ab"`
	KdKepemilikan          string  `json:"kd_kepemilikan"`
	TgKepemilikan          string  `json:"tg_kepemilikan"`
	JamKepemilikan         string  `json:"jam_kepemilikan"`
	KdBlockir              *string `json:"kd_blockir"`
	TgFaktur               *string `json:"tg_faktur"`
	TgFiskal               *string `json:"tg_fiskal"`
	TgGantiMesin           *string `json:"tg_gntmesin"`
	TgKwitansi             *string `json:"tg_kwitansi"`
	TgUben                 *string `json:"tg_uben"`
	TgBlockir              *string `json:"tg_blockir"`
	NilaiJual              string  `json:"nilai_jual"`
	Bobot                  string  `json:"bobot"`
	PabDasar               string  `json:"pab_dasar"`
	NoKTP                  string  `json:"no_ktp"`
	NPWP                   string  `json:"npwp"`
	NmPemilik              string  `json:"nm_pemilik"`
	AlPemilik              string  `json:"al_pemilik"`
	KdWilSubjekPajak       string  `json:"kd_wil_subjek_pajak"`
	KdKelurahan            string  `json:"kd_kelurahan"`
	KdKecamatan            string  `json:"kd_kecamatan"`
	Source                 string  `json:"source"`
	TgUpdate               string  `json:"tg_update"`
	JamUpdate              string  `json:"jam_update"`
	RT                     string  `json:"rt"`
	RW                     string  `json:"rw"`
	IDSubjekPajakType      *string `json:"id_subjek_pajak_type"`
	Email                  *string `json:"email"`
	NoWA                   string  `json:"no_wa"`
	NoHP                   *string `json:"no_hp"`
	KdPos                  *string `json:"kd_pos"`
	Nib                    *string `json:"nib"`
	StatusBayarPotensiBaru string  `json:"status_bayar_potensi_baru"`
}

type HitungObjekPAB struct {
	NoAb1            string  `json:"no_ab1" gorm:"column:no_ab1" form:"no_ab1"`
	NoAb2            string  `json:"no_ab2" gorm:"column:no_ab2" form:"no_ab2"`
	NoAb3            string  `json:"no_ab3" gorm:"column:no_ab3" form:"no_ab3"`
	TgAkhirPajak     string  `json:"tg_akhir_pajak" gorm:"column:tg_akhir_pajak" form:"tg_akhir_pajak"`
	TgAkhirPajakBaru string  `json:"tg_akhir_pajak_baru" gorm:"column:tg_akhir_pajak_baru" form:"tg_akhir_pajak_baru"`
	NilaiJual        string  `json:"nilai_jual" gorm:"column:nilai_jual" form:"nilai_jual"`
	Bobot            string  `json:"bobot" gorm:"column:bobot" form:"bobot"`
	KdKepemilikan    string  `json:"kd_kepemilikan" gorm:"column:kd_kepemilikan" form:"kd_kepemilikan"`
	KdJenisAb        string  `json:"kd_jenis_ab" gorm:"column:kd_jenis_ab" form:"kd_jenis_ab"`
	KdBlockir        *string `json:"kd_blockir" gorm:"column:kd_blockir" form:"kd_blockir"`
	ByrKdpn          string  `json:"byr_kdpn" gorm:"column:byr_kdpn" form:"byr_kdpn"`
	TgDaftar         string  `json:"tg_daftar" gorm:"column:tg_daftar" form:"tg_daftar"`
	KdWil            string  `json:"kd_wil" gorm:"column:kd_wil" form:"kd_wil"`
}

type HasilHitungPab struct {
	BeaBBNABPok int64 `json:"bea_bbnab_pok"`
	BeaBBNABDen int64 `json:"bea_bbnab_den"`
	BeaPABPok   int64 `json:"bea_pab_pok"`
	BeaPABDen   int64 `json:"bea_pab_den"`
	BeaPABPok1  int64 `json:"bea_pab_pok1"`
	BeaPABDen1  int64 `json:"bea_pab_den1"`
	BeaPABPok2  int64 `json:"bea_pab_pok2"`
	BeaPABDen2  int64 `json:"bea_pab_den2"`
	BeaPABPok3  int64 `json:"bea_pab_pok3"`
	BeaPABDen3  int64 `json:"bea_pab_den3"`
	BeaPABPok4  int64 `json:"bea_pab_pok4"`
	BeaPABDen4  int64 `json:"bea_pab_den4"`
	BeaPABPok5  int64 `json:"bea_pab_pok5"`
	BeaPABDen5  int64 `json:"bea_pab_den5"`

	BeaBBNABPokD int64 `json:"bea_bbnab_pok_d"`
	BeaBBNABDenD int64 `json:"bea_bbnab_den_d"`
	BeaPABPokD   int64 `json:"bea_pab_pok_d"`
	BeaPABDenD   int64 `json:"bea_pab_den_d"`
	BeaPABPok1D  int64 `json:"bea_pab_pok1_d"`
	BeaPABDen1D  int64 `json:"bea_pab_den1_d"`
	BeaPABPok2D  int64 `json:"bea_pab_pok2_d"`
	BeaPABDen2D  int64 `json:"bea_pab_den2_d"`
	BeaPABPok3D  int64 `json:"bea_pab_pok3_d"`
	BeaPABDen3D  int64 `json:"bea_pab_den3_d"`
	BeaPABPok4D  int64 `json:"bea_pab_pok4_d"`
	BeaPABDen4D  int64 `json:"bea_pab_den4_d"`
	BeaPABPok5D  int64 `json:"bea_pab_pok5_d"`
	BeaPABDen5D  int64 `json:"bea_pab_den5_d"`

	Ket1             string `json:"ket1"`
	Ket2             string `json:"ket2"`
	Ket3             string `json:"ket3"`
	TgAkhirPajak     string `json:"tg_akhir_pajak"`
	TgAkhirPajakBaru string `json:"tg_akhir_pajak_baru"`
	IsAbleBayar      bool   `json:"is_able_bayar"`
	SelangHari       int    `json:"selang_hari"`
	SelangBulan      int    `json:"selang_bulan"`
	SelangTahun      int    `json:"selang_tahun"`
}

type InfoPabFlat struct {
	IdObjekPajak           string  `json:"id_objek_pajak"`
	NoAb1                  string  `json:"no_ab1"`
	NoAb2                  string  `json:"no_ab2"`
	NoAb3                  string  `json:"no_ab3"`
	KdStatus               string  `json:"kd_status"`
	KdMerekAb              string  `json:"kd_merek_ab"`
	ThBuatan               string  `json:"th_buatan"`
	KdMohon1               string  `json:"kd_mohon1"`
	KdMohon2               string  `json:"kd_mohon2"`
	KdMohon3               string  `json:"kd_mohon3"`
	KdMohon4               string  `json:"kd_mohon4"`
	KdMohon5               string  `json:"kd_mohon5"`
	KdMohon6               string  `json:"kd_mohon6"`
	NmMerekAb              string  `json:"nm_merek_ab"`
	NmModelAb              string  `json:"nm_model_ab"`
	NoRangka               string  `json:"no_rangka"`
	NoMesin                string  `json:"no_mesin"`
	WarnaAb                string  `json:"warna_ab"`
	IdBbm                  string  `json:"id_bbm"`
	JumlahCc               string  `json:"jumlah_cc"`
	IdSubjekPajak          string  `json:"id_subjek_pajak"`
	KdWil                  string  `json:"kd_wil"`
	KdWilProses            string  `json:"kd_wil_proses"`
	KdJenisAB              string  `json:"kd_jenis_ab"`
	KdKepemilikan          string  `json:"kd_kepemilikan"`
	TgKepemilikan          string  `json:"tg_kepemilikan"`
	JamKepemilikan         string  `json:"jam_kepemilikan"`
	KdBlockir              *string `json:"kd_blockir"`
	TgFaktur               *string `json:"tg_faktur"`
	TgFiskal               *string `json:"tg_fiskal"`
	TgGntMesin             *string `json:"tg_gntmesin"`
	TgKwitansi             *string `json:"tg_kwitansi"`
	TgUben                 *string `json:"tg_uben"`
	TgBlockir              *string `json:"tg_blockir"`
	NilaiJual              string  `json:"nilai_jual"`
	Bobot                  string  `json:"bobot"`
	PabDasar               string  `json:"pab_dasar"`
	NoKtp                  string  `json:"no_ktp"`
	Npwp                   string  `json:"npwp"`
	NmPemilik              string  `json:"nm_pemilik"`
	AlPemilik              string  `json:"al_pemilik"`
	KdWilSubjekPajak       string  `json:"kd_wil_subjek_pajak"`
	KdKelurahan            string  `json:"kd_kelurahan"`
	KdKecamatan            string  `json:"kd_kecamatan"`
	Source                 string  `json:"source"`
	TgUpdate               string  `json:"tg_update"`
	JamUpdate              string  `json:"jam_update"`
	Rt                     string  `json:"rt"`
	Rw                     string  `json:"rw"`
	IdSubjekPajakType      *string `json:"id_subjek_pajak_type"`
	Email                  *string `json:"email"`
	NoWa                   string  `json:"no_wa"`
	NoHp                   *string `json:"no_hp"`
	KdPos                  *string `json:"kd_pos"`
	Nib                    *string `json:"nib"`
	StatusBayarPotensiBaru string  `json:"status_bayar_potensi_baru"`

	// Bea dan Denda (asumsi int)
	BeaBBNABPok int `json:"bea_bbnab_pok"`
	BeaBBNABDen int `json:"bea_bbnab_den"`
	BeaPABPok   int `json:"bea_pab_pok"`
	BeaPABDen   int `json:"bea_pab_den"`
	BeaPABPok1  int `json:"bea_pab_pok1"`
	BeaPABDen1  int `json:"bea_pab_den1"`
	BeaPABPok2  int `json:"bea_pab_pok2"`
	BeaPABDen2  int `json:"bea_pab_den2"`
	BeaPABPok3  int `json:"bea_pab_pok3"`
	BeaPABDen3  int `json:"bea_pab_den3"`
	BeaPABPok4  int `json:"bea_pab_pok4"`
	BeaPABDen4  int `json:"bea_pab_den4"`
	BeaPABPok5  int `json:"bea_pab_pok5"`
	BeaPABDen5  int `json:"bea_pab_den5"`

	BeaBBNABPokD int `json:"bea_bbnab_pok_d"`
	BeaBBNABDenD int `json:"bea_bbnab_den_d"`
	BeaPABPokD   int `json:"bea_pab_pok_d"`
	BeaPABDenD   int `json:"bea_pab_den_d"`
	BeaPABPok1D  int `json:"bea_pab_pok1_d"`
	BeaPABDen1D  int `json:"bea_pab_den1_d"`
	BeaPABPok2D  int `json:"bea_pab_pok2_d"`
	BeaPABDen2D  int `json:"bea_pab_den2_d"`
	BeaPABPok3D  int `json:"bea_pab_pok3_d"`
	BeaPABDen3D  int `json:"bea_pab_den3_d"`
	BeaPABPok4D  int `json:"bea_pab_pok4_d"`
	BeaPABDen4D  int `json:"bea_pab_den4_d"`
	BeaPABPok5D  int `json:"bea_pab_pok5_d"`
	BeaPABDen5D  int `json:"bea_pab_den5_d"`

	Ket1 string `json:"ket1"`
	Ket2 string `json:"ket2"`
	Ket3 string `json:"ket3"`

	// Hasil Hitung Pajak
	TgAkhirPajak     string `json:"tg_akhir_pajak"`
	TgAkhirPajakBaru string `json:"tg_akhir_pajak_baru"`
	IsAbleBayar      bool   `json:"is_able_bayar"`
	SelangHari       int    `json:"selang_hari"`
	SelangBulan      int    `json:"selang_bulan"`
	SelangTahun      int    `json:"selang_tahun"`
}

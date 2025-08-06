package dto

type RequestGetSubjek struct {
	Field string `json:"field" form:"field" binding:"required"`
	Value string `json:"value" form:"value" binding:"required"`
}

type ResponseGetSubjek struct {
	IDSubjekPajak     string  `json:"id_subjek_pajak"`
	NoKtp             *string `json:"no_ktp"`
	Npwp              *string `json:"npwp"`
	NmPemilik         string  `json:"nm_pemilik"`
	AlPemilik         string  `json:"al_pemilik"`
	KdWilSubjekPajak  string  `json:"kd_wil_subjek_pajak"`
	KdKelurahan       string  `json:"kd_kelurahan"`
	KdKecamatan       string  `json:"kd_kecamatan"`
	Source            string  `json:"source"`
	TgUpdate          string  `json:"tg_update"`
	JamUpdate         string  `json:"jam_update"`
	RT                string  `json:"rt"`
	RW                string  `json:"rw"`
	IDSubjekPajakType string  `json:"id_subjek_pajak_type"`
	Email             *string `json:"email"`
	NoWa              *string `json:"no_wa"`
	NoHp              *string `json:"no_hp"`
	KdPos             string  `json:"kd_pos"`
	Nib               *string `json:"nib"`
	Npwpd             *string `json:"npwpd"`
	UpdatedAt         string  `json:"updated_at"`
}

type GetKepemilikanRequest struct {
	IdSubjekPajak string `json:"id_subjek_pajak" form:"id_subjek_pajak"`
}

type GetKepemilikanAbResponse struct {
	NoAb1        string `json:"no_Ab1" gorm:"column:no_ab1"`
	NoAb2        string `json:"no_Ab2" gorm:"column:no_ab2"`
	NoAb3        string `json:"no_Ab3" gorm:"column:no_ab3"`
	KdWil        string `json:"kd_wil" gorm:"column:kd_wil"`
	NmWil        string `json:"nm_wil" gorm:"column:nm_wil"`
	NmMerekAB    string `json:"nm_merek_ab" gorm:"column:nm_merek_ab"`
	NmModelAB    string `json:"nm_model_ab" gorm:"column:nm_model_ab"`
	ThBuatan     string `json:"th_buatan" gorm:"column:th_buatan"`
	NmJenisAB    string `json:"nm_jenis_ab" gorm:"column:nm_jenis_ab"`
	TgAkhirPajak string `json:"tg_akhir_pajak" gorm:"column:tg_akhir_pajak"`
}
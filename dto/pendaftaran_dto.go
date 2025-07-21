package dto

type CekObjekRequest struct {
	NoAb1 string `form:"no_ab1" json:"no_ab1" binding:"required"`
	NoAb2 string `form:"no_ab2" json:"no_ab2" binding:"required"`
	NoAb3 string `form:"no_ab3" json:"no_ab3" binding:"required"`
}

type CekTransaksiResponse struct {
	IsExist     bool   `json:"is_exist"`
	IdTrnab     string `json:"id_trnab"`
	KdStatus    string `json:"kd_status"`
	KdWilProses string `json:"kd_wil_proses"`
	TgDaftar    string `json:"tg_daftar"`
	JamDaftar   string `json:"jam_daftar"`
	TgTetap     string `json:"tg_tetap"`
	JamTetap    string `json:"jam_tetap"`
	TgBayar     string `json:"tg_bayar"`
	JamBayar    string `json:"jam_bayar"`
}

type PendaftaranTahunanRequest struct {
	NoAb1 string `json:"no_ab1" binding:"required"`
	NoAb2 string `json:"no_ab2" binding:"required"`
	NoAb3 string `json:"no_ab3" binding:"required"`
	IdUser string `json:"id_user" binding:"required"`
}

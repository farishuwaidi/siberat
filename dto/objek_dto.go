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
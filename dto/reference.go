package dto

import (
	"encoding/json"
	"time"
)

type ResponseKodeRole struct {
	IDRole   string `json:"id_role" gorm:"column:id_role"` 
	NameRole string `json:"nm_role" gorm:"column:nm_role"`
}

type ResponseKodeRoleParse struct {
	IDRole   int `json:"id_role" gorm:"column:id_role"` // Role ID (string → dikonversi ke int)
	NameRole string `json:"nm_role" gorm:"column:nm_role"`
}

type ResponseKodeWil struct {
	KdWil    string `json:"kd_wil" gorm:"column:kd_wil"`     // Code of the region
	NmWil    string `json:"nm_wil" gorm:"column:nm_wil"`     // Name of the region
	AlWil    string `json:"al_wil" gorm:"column:al_wil"`     // Address of the region
	KabKota  string `json:"kab_kota" gorm:"column:kab_kota"` // City or district of the region
	Provinsi string `json:"provinsi" gorm:"column:provinsi"` // Province of the region
}

type ResponseKdJenisAB struct {
	KdJenisAB string `json:"kd_jenis_ab" gorm:"kd_jenis_ab"`
	NmJenisAB string `json:"nm_jenis_ab" gorm:"nm_jenis_ab"`
}

type RequestKdMerek struct {
	KdJenisAB string `json:"kd_jenis_ab" form:"kd_jenis_ab"`
}

type ResponseKdMerek struct {
	KdMerekAB string `json:"kd_merek_ab" gorm:"column:kd_merek_ab"`
	NmMerekAB string `json:"nm_merek_ab" gorm:"column:nm_merek_ab"`
}

type RequestKdModel struct {
	KdMerekAB string `json:"kd_merek_ab" form:"kd_merek_ab"`
}

type ResponseKdModel struct {
	KdMerekAB  string `json:"kd_merek_ab" gorm:"column:kd_merek_ab"`
	NmMerekAB  string `json:"nm_merek_ab" gorm:"column:nm_merek_ab"`
	NmModelkAB string `json:"nm_model_ab" gorm:"column:nm_model_ab"`
}

type RequestNilaiJual struct {
	KdMerekAB string `json:"kd_merek_ab" form:"kd_merek_ab"`
}

type ResponseNilaiJual struct {
	KdMerekAB string  `json:"kd_merek_ab" gorm:"column:kd_merek_ab"`
	ThBuatan  int     `json:"th_buatan" gorm:"column:th_buatan"`
	NilaiJual string  `json:"nilai_jual" gorm:"column:nilai_jual"`
	Bobot     string  `json:"bobot" gorm:"column:bobot"`
	PabDasar  string  `json:"pab_dasar" gorm:"column:pab_dasar"`
	EditBY    *string `json:"edit_by" gorm:"column:edit_by"`
	UpdatedAt *time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

// Custom MarshalJSON untuk memformat UpdatedAt
func (r ResponseNilaiJual) MarshalJSON() ([]byte, error) {
	type Alias ResponseNilaiJual
	return json.Marshal(&struct {
		UpdatedAt *string `json:"updated_at"`
		Alias
	}{
		UpdatedAt: func() *string {
			if r.UpdatedAt != nil {
				str := r.UpdatedAt.Format("2006-01-02 15:04:05")
				return &str
			}
			return nil
		}(),
		Alias: (Alias)(r),
	})
}

type ResponseJenisBBM struct {
	Id string `json:"id" gorm:"column:id"`
	NmBbm string `json:"nm_bbm" gorm:"column:nm_bbm"`
}

type ResponseJenisKepemilikan struct {
	KdKepemilikan string `json:"kd_kepemilikan"`
	NmKepemilikan string `json:"nm_kepemilikan"`
}

type ResponseJenisKepemilikanParse struct {
	KdKepemilikan int `json:"kd_kepemilikan"`
	NmKepemilikan string `json:"nm_kepemilikan"` 
}

type ResponseKdMohon map[string]interface{}
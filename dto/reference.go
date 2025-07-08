package dto

type ResponseKodeRole struct {
	IDRole   string `json:"id_role" gorm:"column:id_role"` // Role ID (string → dikonversi ke int)
	NameRole string `json:"nm_role" gorm:"column:nm_role"` // Description of the role
}

type ResponseKodeWil struct {
	KdWil    string `json:"kd_wil" gorm:"column:kd_wil"`     // Code of the region
	NmWil    string `json:"nm_wil" gorm:"column:nm_wil"`     // Name of the region
	AlWil    string `json:"al_wil" gorm:"column:al_wil"`     // Address of the region
	KabKota  string `json:"kab_kota" gorm:"column:kab_kota"` // City or district of the region
	Provinsi string `json:"provinsi" gorm:"column:provinsi"` // Province of the region
}
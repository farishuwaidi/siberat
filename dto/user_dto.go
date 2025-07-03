package dto

type UpdateUserRequest struct {
	Name        string `json:"name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	NoHp        string `json:"noHp"`
	IDRole      string `json:"id_role"` // Role ID (string → dikonversi ke int)
	KodeWilayah string `json:"kodeWilayah"`
}
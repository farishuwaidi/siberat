package dto

type UpdateUserRequest struct {
	Name        string `json:"name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	NoHp        string `json:"noHp"`
	Role        string `json:"role"` // Role ID (string → dikonversi ke int)
	KodeWilayah string `json:"kodeWilayah"`
}
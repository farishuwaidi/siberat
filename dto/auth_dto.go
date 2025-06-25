package dto

type RegisterRequest struct {
	Name        string `json:"name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	NoHp        string `json:"noHp"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	KodeWilayah string `json:"kodeWilayah"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Role  string `json:"role"`
}

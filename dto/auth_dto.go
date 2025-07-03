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
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID                  int     `json:"id"`
	Name                string  `json:"sur_name"`
	Email               string  `json:"email"`
	UserName            string  `json:"username"`
	EmailVerifiedAt     *string `json:"email_verified_at"`
	UserData            string  `json:"user_data"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
	KodeWilayah         string  `json:"kd_wil"`
	KodeWilayahProses   string  `json:"kd_wil_kerja"`
	NoHp                string  `json:"no_hp"`
	NoWa                string  `json:"no_wa"`
	UserNameSipandu     string  `json:"username_sipandu"`
	UserPasswordSipandu string  `json:"userpassword_sipandu"`
	IDRole              int     `json:"id_role"`
	Token               string  `json:"token"`
	TokenExpiredAt      string  `json:"token_expired_at"` // Optional, if you want to include token expiration
	PhotoUrl            string  `json:"photo_url"`        // Optional, if you want to include user's photo URL
}

type LoginResponseParam struct {
	Code    string        `json:"code"`
	Success bool          `json:"success"`
	Data    LoginResponse `json:"data"`
	Message string        `json:"message"`
	Param   LoginRequest  `json:"param"`
}

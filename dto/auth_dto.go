package dto

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

type GetPermissionRequest struct {
	Id string `form:"id" binding:"required"` // ID of the user
}

type GetPermissionResponse struct {
	IdRole              string `json:"id_role"`                 // Role ID of the user
	NmRole              string `json:"nm_role"`                 // Name of the role
	StsDaftar           string `json:"sts_daftar"`              // Registration status of the user
	StsTambahPotensi    string `json:"sts_tambah_potensi"`      // Status of adding potential
	StsInfoPab          string `json:"sts_info_pab"`            // Status of PAB information
	StsKirimUlgSkkp     string `json:"sts_kirim_ulg_skkp"`      // Status of sending SKKP
	StsCetakNab         string `json:"sts_cetak_nab"`           // Status of printing NAB
	StsCetakUlgNab      string `json:"sts_cetak_ulg_nab"`       // Status of reprinting NAB
	StsInfoSubjekPab    string `json:"sts_info_subjek_pab"`     // Status of PAB subject information
	StsDaftarUlgPotBaru string `json:"sts_daftar_ulg_pot_baru"` // Status of registering new potential
	StsUserMngmt        string `json:"sts_user_mngmt"`          // User management status
	StsNjabMngmt        string `json:"sts_njab_mngmt"`          // NJAB management status
	StsInfoNjab         string `json:"sts_info_njab"`           // Information status of NJAB
	StsInfoUser         string `json:"sts_info_user"`           // User information status
	StsTable            string `json:"sts_table"`               // Table status
	StsTableMngmt       string `json:"sts_table_mngmt"`         // Table management status
	StsVerif            string `json:"sts_verif"`               // Verification status
}

type UpdatePassword struct {
	ID              int    `json:"id"`
	Password        string `json:"password"`
	PasswordBaru    string `json:"password_baru"`
	PasswordConfirm string `json:"password_confirm"`
}

type UpdatePasswordResponse struct {
	Message string `json:"message"`
}
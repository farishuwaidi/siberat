package entity

import "time"

type User struct {
	ID        	int `gorm:"column:id"`
	Name      	string `gorm:"column:sur_name"`
	Email     	string `gorm:"column:email"`
	UserName 	string `gorm:"column:username"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at"`
	Password  	string `gorm:"column:password"`
	UserData 	string `gorm:"column:user_data"`
	RememberToken string `gorm:"column:remember_token"`
	CreatedAt 	time.Time `gorm:"column:created_at"`
	UpdatedAt 	time.Time	`gorm:"column:updated_at"`
	KodeWilayah string `gorm:"column:kd_wil"`
	KodeWilayahProses string `gorm:"column:kd_wil_kerja"`
	NoHp 		string `gorm:"column:no_hp"`
	NoWa 		string `gorm:"column:no_wa"`
	UserNameSipandu string `gorm:"column:username_sipandu" default:"sistem_pab"`
	UserPasswordSipandu string `gorm:"column:userpassword_sipandu" default:"1"`
	IDRole 		int `gorm:"column:id_role"`
}
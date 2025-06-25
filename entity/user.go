package entity

import "time"

type User struct {
	ID        	int
	Name      	string
	UserName 	string
	Email     	string
	NoHp 		string
	RoleID 		int
	Role 		RoleUser `gorm:"foreignKey:RoleID"`
	Password  	string
	KodeWilayah string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}
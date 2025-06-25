package entity

type RoleUser struct {
	ID   int    `gorm:"primarykey"`
	Role string `gorm:"unique"`
}
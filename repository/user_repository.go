package repository

import (
	"Siberat/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository{
	return	&userRepository{db}
}

func(r *userRepository) GetAllUser()([]entity.User, error) {
	var users []entity.User
	err := r.db.Preload("Role").Find(&users).Error
	return users, err
}
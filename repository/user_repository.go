package repository

import (
	"Siberat/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]entity.User, error)
	UpdateUser(id int, updatedUser entity.User) error
	GetUserByID(id int) (*entity.User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUser() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Preload("Role").Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUser(id int, updatedUser entity.User) error {
	return r.db.Model(&entity.User{}).Where("id = ?", id).Updates(updatedUser).Error
}

func (r *userRepository) GetUserByID(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("Role").First(&user, id).Error
	return &user, err
}

func (r *userRepository) DeleteUser(id int) error {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&user).Error
}

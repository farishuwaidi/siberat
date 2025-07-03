package repository

import (
	"Siberat/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user *entity.User) error
	EmailExist(email string) bool
	GetUserByUsername(username string) (*entity.User, error)
	PreLoadUserRole(user *entity.User)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Register(user *entity.User) error {
	return r.db.Create(&user).Error
}

func (r *authRepository) EmailExist(email string) bool {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error
	return err == nil
}

func (r *authRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "username = ?", username).First(&user).Error
	return &user, err
}

func (r *authRepository) PreLoadUserRole(user *entity.User) {
	r.db.Preload("Role").First(user, user.ID)
}

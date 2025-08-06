package repository

import (
	"Siberat/entity"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(email string) bool
	GetUserByUsername(username string) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	GetRoleByID(id int) (*entity.RoleUser, error)
	UpdatePassword(id int, newHashPassword string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
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

func (r *authRepository) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id= ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetRoleByID(id int) (*entity.RoleUser, error) {
	var role entity.RoleUser
	idstr := fmt.Sprintf("%d", id)
	if err := r.db.
		Raw("SELECT * FROM t_role WHERE id_role = ?", idstr).
		Scan(&role).Error;err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *authRepository) UpdatePassword(id int, newHashPassword string) error {
	return r.db.Model(&entity.User{}).
		Where("id = ? ", id).
		Update("password", newHashPassword).Error
}
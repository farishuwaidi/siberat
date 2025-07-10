package repository

import (
	"Siberat/dto"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetPaginatedUsers(filter string, pqge, perPage int)([]dto.UserDataItem, int64, error)
	GetUserDetail(id string)(*dto.GetUserDetailResponse, error)
	AddUserData(data map[string]interface{}) error
	EditUserData(id any, data map[string]interface{}) error
	DeleteUserData(id any) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetPaginatedUsers(filter string, page, perPage int) ([]dto.UserDataItem, int64, error) {
	var users []dto.UserDataItem
	var total int64

	offset := (page - 1) * perPage
	like := "%" + filter + "%"

	query := r.db.Table("users AS u").
		Select(`u.id, u.sur_name, u.username, wil.nm_wil, u.kd_wil, u.kd_wil_kerja, trole.nm_role`).
		Joins("LEFT JOIN t_role AS trole ON trole.id_role = u.id_role::text").
		Joins("LEFT JOIN t_wiluppd AS wil ON wil.kd_wil = u.kd_wil")
		
	fmt.Println(query.Statement.SQL.String())

	if filter != "" {
		query = query.Where(`(
			CAST(u.id AS TEXT) ILIKE ? OR
			u.sur_name ILIKE ? OR
			u.username ILIKE ? OR
			wil.nm_wil ILIKE ? OR
			trole.nm_role ILIKE ?
		)`, like, like, like, like, like)
	}

	if err := query.Limit(perPage).Offset(offset).Scan(&users).Error; err != nil {
		return nil, 0, err
	}

	countQuery := r.db.Table("users AS u").
		Joins("LEFT JOIN t_role AS trole ON trole.id_role = u.id_role::text").
		Joins("LEFT JOIN t_wiluppd AS wil ON wil.kd_wil = u.kd_wil")

	if filter != "" {
		countQuery = countQuery.Where(`(
			CAST(u.id AS TEXT) ILIKE ? OR
			u.sur_name ILIKE ? OR
			u.username ILIKE ? OR
			wil.nm_wil ILIKE ? OR
			trole.nm_role ILIKE ?
		)`, like, like, like, like, like)
	}

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	for i := range users {
		users[i].PhotoUrl = "https://ui-avatars.com/api/"
	}

	return users, total, nil
}

func (r *userRepository) GetUserDetail(id string) (*dto.GetUserDetailResponse, error) {
	var user dto.GetUserDetailResponse
	err := r.db.
		Table("users").
		Select(`id, sur_name, email, username, email_verified_at, user_data, created_at, updated_at,
				kd_wil, kd_wil_kerja, no_hp, no_wa, username_sipandu, userpassword_sipandu, id_role`).
		Where("id = ?", id).
		Scan(&user).Error
	
	if err != nil {
		return nil, err
	}

	user.PhotoUrl = "https://ui-avatars.com/api/"
	return &user, nil
}



func (r *userRepository) AddUserData(data map[string]interface{}) error {
	return r.db.Table("users").Create(data).Error
}

func (r *userRepository) EditUserData(id any, data map[string]interface{}) error {
	return r.db.Table("users").Where("id = ?", id).Updates(data).Error
}

func (r *userRepository) DeleteUserData(id any) error {
	return r.db.Table("users").Where("id = ?", id).Delete(nil).Error
}

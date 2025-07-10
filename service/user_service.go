package service

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/repository"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type UserService interface {
	GetUserData(req dto.GetUserDataRequest) (*dto.GetUserDataResponse, error)
	GetUserDetail(id string) (*dto.GetUserDetailResponse, error)
	UpdateUserData(req dto.UpdateUserDataRequest) (any, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		repository: r,
	}
}

func (s *userService) GetUserData(req dto.GetUserDataRequest) (*dto.GetUserDataResponse, error) {
	page ,err := strconv.Atoi(req.Page)
	if err != nil || page<= 0{
		page =1
	}
	perPage, err := strconv.Atoi(req.PerPage)
	if err != nil || perPage <=0 {
		perPage = 20
	}
	users, total, err := s.repository.GetPaginatedUsers(req.FilterBy, page, perPage)
	if err != nil {
		return nil, err
	}

	lasPage := (int(total) + perPage - 1)/perPage
	urlPath := "http://localhost:8080/api/user/get_user"


	res := &dto.GetUserDataResponse{
		CurrentPage: page,
		Data: users,
		FirstPageUrl: fmt.Sprintf("%s?page=1", urlPath),
		From: (page-1)*perPage+1,
		LastPage: lasPage,
		LastPageUrl: fmt.Sprintf("%s?page=%d", urlPath,lasPage),
		Links: dto.GeneratePaginateLinks(urlPath,page, lasPage),
		NextPageUrl: nil,
		Path: urlPath,
		PerPage: strconv.Itoa(perPage),
		PrevPageUrl: nil,
		To: (page-1)*perPage+len(users),
		Total: total,
	}
	if page > 1 {
		prev :=fmt.Sprintf("%s?page=%d", urlPath, page-1)
		res.PrevPageUrl = &prev
	}
	if page < lasPage {
		next := fmt.Sprintf("%s?page=%d", urlPath, page+1)
		res.NextPageUrl = &next
	}
	return res, nil
}

func (s *userService) GetUserDetail(id string) (*dto.GetUserDetailResponse, error) {
	return s.repository.GetUserDetail(id)
}

func (s *userService) UpdateUserData(req dto.UpdateUserDataRequest) (any, error) {
	switch req.Method {
	case "A":
		uid := helper.GenerateHexID()
		data := req.DataCrud

		newUser := map[string]interface{}{
			"id": uid,
			"sur_name": data["sur_name"],
			"email": data["email"],
			"username": data["username"],
			"email_verified_at": time.Now(),
			"password": helper.HashPassword(data["password"].(string)),
			"created_at": time.Now(),
			"updated_at": time.Now(),
			"kd_wil": data["kd_wil"],
			"kd_wil_kerja": data["kd_wil_kerja"],
			"no_hp": data["no_hp"],
			"no_wa": data["no_wa"],
			"username_sipandu": "sistem_pab",
			"userpassword_sipandu": "1",
			"id_role": data["id_role"],
		}

		if err := s.repository.AddUserData(newUser); err != nil {
			return nil, err
		}
		return s.repository.GetUserDetail(fmt.Sprint(uid))
	case "E":
		idEdited := req.DataCrud["id"]
		dataEdit := req.DataCrud["data_edit"].(map[string]any)
		if pw, ok := dataEdit["password"]; ok {
			dataEdit["password"] = helper.HashPassword(pw.(string))
		}
		if err := s.repository.EditUserData(idEdited, dataEdit); err != nil {
			return nil, err
		}

		idStr, ok := idEdited.(string)
		if !ok {
			return nil, errors.New("idEdited is not string")
		}
		return s.repository.GetUserDetail(idStr)
	
	case "D":
		idDel := req.DataCrud["id"]
		if err := s.repository.DeleteUserData(idDel); err != nil {
			return nil, err
		}
		return req, nil
	}

	return nil, errors.New("method not found")

}

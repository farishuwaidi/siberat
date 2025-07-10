package dto

import "fmt"

type GetUserDataRequest struct {
	IdUser   string `form:"id" json:"id_user" binding:"required"`
	Page     string `form:"page" json:"page" binding:"required"`
	PerPage  string `form:"per_page" json:"per_page" binding:"required"`
	FilterBy string `form:"filter_by" json:"filter_by"`
}

type UserDataItem struct {
	ID         int    `json:"id"`
	SurName    string `json:"sur_name"`
	Username   string `json:"username"`
	NmWil      string `json:"nm_wil"`
	KdWil      string `json:"kd_wil"`
	KdWilKerja string `json:"kd_wil_kerja"`
	NmRole     string `json:"nm_role"`
	PhotoUrl   string `json:"photo_url"`
}

type PaginationLink struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}

type GetUserDataResponse struct {
	CurrentPage  int              `json:"current_page"`
	Data         []UserDataItem   `json:"data"`
	FirstPageUrl string           `json:"first_page_url"`
	From         int              `json:"from"`
	LastPage     int              `json:"last_page"`
	LastPageUrl  string           `json:"last_page_url"`
	Links        []PaginationLink `json:"links"`
	NextPageUrl  *string          `json:"next_page_url"`
	Path         string           `json:"path"`
	PerPage      string           `json:"per_page"`
	PrevPageUrl  *string          `json:"prev_page_url"`
	To           int              `json:"to"`
	Total        int64            `json:"total"`
}

type GetUserDetailRequest struct {
	ID string `form:"id" json:"id" binding:"required"`
}

type GetUserDetailResponse struct {
	ID                  int     `json:"id"`
	SurName             string  `json:"sur_name"`
	Email               string  `json:"email"`
	Username            string  `json:"username"`
	EmailVerifiedAt     *string `json:"email_verified_at"`
	UserData            *string `json:"user_data"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
	KdWil         		string  `json:"kd_wil"`
	KdWilKerja   		string  `json:"kd_wil_kerja"`
	NoHp                string  `json:"no_hp"`
	NoWa                string  `json:"no_wa"`
	UsernameSipandu     string  `json:"username_sipandu"`
	UserpasswordSipandu string  `json:"userpassword_sipandu"`
	IDRole              int     `json:"id_role"`
	PhotoUrl            string  `json:"photo_url"`        // Optional, if you want to include user's photo URL
}

func GeneratePaginateLinks(urlPath string, current, last int) []PaginationLink {
	links := []PaginationLink{
		{URL: nil, Label: "<< Previous", Active: false},
	}

	for i := 1; i <= last; i++ {
		url := fmt.Sprintf("%s?page=%d", urlPath, i)
		links = append(links, PaginationLink{
			URL: &url,
			Label: fmt.Sprintf("%d",i),
			Active: i == current,
		})
	}
	links = append(links, PaginationLink{URL:nil, Label: "Next >>", Active: false})
	return links
}

type UpdateUserDataRequest struct {
	IdUser string `json:"id_user" binding:"required"`
	Method string `json:"method" binding:"required"`
	DataCrud map[string]any `json:"data_crud" binding:"required"`
}
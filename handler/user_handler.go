package handler

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetAllUser(c *gin.Context){
	users, err := h.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			StatusCode:  http.StatusInternalServerError,
			Message: "failed to retrieve users",
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message: "List of users",
		Data: users,
	}))
}
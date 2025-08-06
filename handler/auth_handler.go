package handler

import (
	"Siberat/dto"
	errorhandler "Siberat/errorHandler"
	"Siberat/helper"
	"Siberat/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{service: s}
}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{
			Code: "0002",
			MessageText: "invalid request format",
		}, login)
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.UnAuthorizedError{
			CodeErr: "0032",
			MessageText: "kredensial tidak sesuai",
		}, login)
		return
	}

	fullresponse := dto.LoginResponseParam{
		Code:    "0000",
		Success: true,
		Data:   *result,
		Message: "Login success",
		Param: login,
	}

	c.JSON(http.StatusOK, fullresponse)
}

func (h *authHandler) GetPermissionData(c *gin.Context) {
	var req dto.GetPermissionRequest
	if err := c.ShouldBind(&req); err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{
			Code: "0002",
			MessageText: "invalid request format",
		}, req)
		return
	}

	result, err := h.service.GetPermissionData(req.Id)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.NotFoundError{
			Code: "0003",
			MessageText: "error, tidak ditemukan data dengan data user berdasarkan id",
		},req)
		return
	}

	permission := helper.ResponseWithData{
		Code:   "0000",
		Success: true,
		Data: result,
		Message: "sukses",
		Param: gin.H{"id":req.Id},
	}
	c.JSON (http.StatusOK, permission)
}

func (h *authHandler) UpdatePassword(c *gin.Context) {
	var req dto.UpdatePassword
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "validation fail",
			Data: []string{},
			Param: req,
		}))
		return
	}
	result, err := h.service.UpdatePassword(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code:   http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
			Data:   []string{},
			Param:  req,
		}))
		return
	}
	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code:   http.StatusOK,
		Success: true,
		Message: "Success",
		Data:   result,
		Param:  req,
	}))
}
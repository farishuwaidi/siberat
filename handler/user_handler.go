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

func (h *UserHandler) GetUserData(c *gin.Context){
	var req dto.GetUserDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": "0001",
			"success": false,
			"data": []any{},
			"message": "invalid request",
			"param": req,
		})
		return
	}
	result, err := h.service.GetUserData(req)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0500",
			"success": false,
			"data": []any{},
			"message": err.Error(),
			"param": req,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "0000",
		"success": true,
		"data": result,
		"message": "sukses",
		"param": req,
	})
 }

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	var req dto.GetUserDetailRequest

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

	user, err := h.service.GetUserDetail(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			Code: http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
			Data: []string{},
			Param: req,
		}))
		return
	}

	if user.ID == nil {
		c.JSON(http.StatusNotFound, helper.Response(dto.ResponseParams{
			Code: http.StatusNotFound,
			Success: false,
			Message: "Data not found",
			Data: []string{},
			Param: req,
		}))
		return
	}
	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "Sukses",
		Data: user,
		Param: req,

	}))
}

func (h *UserHandler) UpdateUserData(c *gin.Context) {
	var req dto.UpdateUserDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
			Code: 401,
			Success: false,
			Message: "invalid request",
			Data: []any{},
			Param: req,
		}))
		return
	}

	result, err := h.service.UpdateUserData(req)
	if err != nil {
		c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
			Code: 404,
			Success: false,
			Message: err.Error(),
			Data: []any{},
			Param: req,
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: 200,
		Success: true,
		Message: "Success",
		Data: result,
		Param: req,
	}))
}
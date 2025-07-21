package handler

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PenetapanHandler struct {
	service service.PenetapanService
}

func NewPenetapanHandler(s service.PenetapanService) *PenetapanHandler {
	return &PenetapanHandler{service: s}	
}

func (h *PenetapanHandler) HitungUlang(c *gin.Context) {
	var req dto.HitungUlangRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: "invalid params",
			Data: []string{},
			Param: req,
		}))
		return
	}

	result, err := h.service.HitungUlang(req)
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

	// if result == nil {
	// 	c.JSON(http.StatusNotFound, helper.Response(dto.ResponseParams{
	// 		Code: http.StatusNotFound,
	// 		Success: false,
	// 		Message: "data not found",
	// 		Data: []string{},
	// 		Param: req,
	// 	}))
	// 	return
	// }
	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "sukses",
		Data: result,
		Param: req,
	}))
}

func (h *PenetapanHandler) GetPenetapan(c *gin.Context) {
	var req dto.GetTrnabRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: "invalid param",
			Data: []string{},
			Param: req,
		}))
		return
	}
	result, err := h.service.GetPenetapan(req.IdTrnab)
	if err != nil {
		if err.Error() == "invalid param" {
			c.JSON(http.StatusBadRequest, dto.ResponseParams{
				Code: 400,
				Success: false,
				Message: err.Error(),
			})
			return
		}
		if err.Error() == "data not found" {
			c.JSON(http.StatusNotFound, dto.ResponseParams{
				Code: 404,
				Success: false,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
				Code: http.StatusInternalServerError,
				Success: false,
				Message: err.Error(),
				Data: []string{},
				Param: req,
		}))
		return
	}

	if result == nil {
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
		Message: "sukses",
		Data: result,
		Param: req,
	}))
}
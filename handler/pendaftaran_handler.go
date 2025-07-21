package handler

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PendaftaranHandler struct {
	service service.PendaftaranService
}

func NewPendaftaranHandler(service service.PendaftaranService) *PendaftaranHandler {
	return &PendaftaranHandler{service}
}

func (h *PendaftaranHandler) CekTransaksi(c *gin.Context) {
	var req dto.CekObjekRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: "invalid param",
			Data: []string{},
			Param: req,
		}))
		return
	}

	result, err := h.service.CekTransaksi(req.NoAb1, req.NoAb2, req.NoAb3)
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

	if result == nil {
		c.JSON(http.StatusNotFound, helper.Response(dto.ResponseParams{
			Code: http.StatusNotFound,
			Success: false,
			Message: "data not found",
			Data: []string{},
			Param: req,
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "success",
		Data: result,
		Param: req,
	}))
}

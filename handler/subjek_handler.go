package handler

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SubjekHandler struct {
	service service.SubjekService
}

func NewSubjekHandler(s service.SubjekService) *SubjekHandler {
	return &SubjekHandler{service: s}
}

func (h *SubjekHandler) GetSubjekPajak(c *gin.Context) {
	var req dto.RequestGetSubjek
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

	result, err := h.service.GetSubjekPajak(req.Field, req.Value)
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
			Message: "data tidak ditemukan",
			Data: []string{},
			Param: req,
		}))
		return
	}

	result.IDSubjekPajak = strings.TrimSpace(strings.TrimSpace(result.IDSubjekPajak))
	result.KdKecamatan = strings.TrimSpace(strings.TrimSpace(result.KdKecamatan))
	result.RT = strings.TrimSpace(strings.TrimSpace(result.RT))
	result.RW = strings.TrimSpace(strings.TrimSpace(result.RW))
	if result.NoKtp != nil {
		trimmed := strings.TrimSpace(*result.NoKtp)
		result.NoKtp = &trimmed
	}
	if result.Npwp != nil {
		trimmed := strings.TrimSpace(*result.Npwp)
		result.Npwp = &trimmed
	}
	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "sukses",
		Data: result,
		Param: req,
	}))
}

func (h *SubjekHandler) GetKepemilikanAb(c *gin.Context) {
	var req dto.GetKepemilikanRequest

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
	 result, err := h.service.GetKepemilikanAb(req.IdSubjekPajak)
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
			Message: "data tidak ditemukan",
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
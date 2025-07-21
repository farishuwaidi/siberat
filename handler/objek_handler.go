package handler

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ObjekHandler struct {
	service  service.ObjekService
}

func NewObjekHandler(service service.ObjekService) *ObjekHandler {
	return &ObjekHandler{service}
}

func (h *ObjekHandler) GetObjekByNab(c *gin.Context) {
	var req dto.ObjekPabRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: "parameter not valid",
			Data: []string{},
			Param: req,
		}))
		return
	}

	objek, err := h.service.GetObjekByNab(req.NoAb1, req.NoAb2, req.NoAb3)
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

	if objek == nil {
		c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
			Code: http.StatusNotFound,
			Success: false,
			Message: "Data tidak ditemukan",
			Data: []string{},
			Param: req,
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "sukses",
		Data: objek,
		Param: req,
	}))
}

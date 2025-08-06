package handler

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NjabHandler struct {
	service service.NjabService
}


func NewNjabHandler(service service.NjabService) *NjabHandler {
	return &NjabHandler{service}
}

func (h *NjabHandler) UpdateMerek(c *gin.Context) {
	var req dto.UpdateNjabRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code:   http.StatusBadRequest,
			Success: false,
			Message: "validation fail",
			Data: []string{},
			Param: req,
		}))
		return
	}

	result, err := h.service.UpdateMerek(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			Code:   http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
			Data:   []string{},
			Param:  req,
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code:  http.StatusOK,
		Success: true,
		Message: "Success",
		Data:   result,
		Param:  req,
	}))
}

func (h *NjabHandler) GetMerekDetails(c *gin.Context) {
	var req dto.GetNjabDetailsRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code:   http.StatusBadRequest,
			Success: false,
			Message: "validation fail",
			Data: []string{},
			Param: req,
		}))
		return
	}
	result, err := h.service.GetMerekDetails(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			Code:   http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
			Data:   []string{},
			Param:  req,
		}))
		return
	}
	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code:  http.StatusOK,
		Success: true,
		Message: "Success",
		Data:   result,
		Param:  req,
	}))
}

func (h *NjabHandler) UpdateNilaiJual(c *gin.Context) {
	var req dto.UpdateNjabRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code:   http.StatusBadRequest,
			Success: false,
			Message: "validation fail",
			Data: []string{},
			Param: req,
		}))
		return
	}

	result, err := h.service.UpdateNilaiJual(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			Code:   http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
			Data:   []string{},
			Param:  req,
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code:  http.StatusOK,
		Success: true,
		Message: "Success",
		Data:   result,
		Param:  req,
	}))
}

func (h *NjabHandler) GetDetailNjab(c *gin.Context) {
	var req dto.GetNilaiJualRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code:   http.StatusBadRequest,
			Success: false,
			Message: "validation fail",
			Data: []string{},
			Param: req,
		}))
		return
	}
	result, err := h.service.GetDetailNjab(req.KdMerekAb, req.ThBuatan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			Code:   http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
			Data:   []string{},
			Param:  req,
		}))
		return
	}
	if result == nil {
		c.JSON(http.StatusNotFound, helper.Response(dto.ResponseParams{
			Code:   http.StatusNotFound,
			Success: false,
			Message: "data not found",
			Data:   []string{},
			Param:  req,
		}))
		return
	}
	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code:  http.StatusOK,
		Success: true,
		Message: "Success",
		Data:   result,
		Param:  req,
	}))
}

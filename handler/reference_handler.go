package handler

import (
	"Siberat/dto"
	"Siberat/helper"
	"Siberat/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReferenceHandler struct {
	service service.ReferenceService
}

func NewReferenceHandler(s service.ReferenceService) *ReferenceHandler {
	return &ReferenceHandler{service: s}
}

func (h *ReferenceHandler) GetKodeRole(c *gin.Context) {
	role, err := h.service.GetKodeRole()
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "success": false, "error": err.Error()})
		return
	}

	var result []dto.ResponseKodeRoleParse 
	for _, i := range role {
		idInt,_ := strconv.Atoi(strings.TrimSpace(i.IDRole))
		result = append(result, dto.ResponseKodeRoleParse{
			IDRole: idInt,
			NameRole: i.NameRole,
		})
	}
	c.JSON(200, gin.H{
		"code": 200, 
		"success": true, 
		"data": result, 
		"message": "sukses", 
		"param": gin.H{
			"query": nil,
		}})
}

func (h *ReferenceHandler) GetKodeWil(c *gin.Context) {
	wilayah, err := h.service.GetKodeWil()
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "success": false, "error": err.Error()})
		return
	}

	for i := range wilayah {
		wilayah[i].KdWil = strings.TrimSpace(wilayah[i].KdWil)
		wilayah[i].NmWil = strings.TrimSpace(wilayah[i].NmWil)
		wilayah[i].AlWil = strings.TrimSpace(wilayah[i].AlWil)
		wilayah[i].KabKota = strings.TrimSpace(wilayah[i].KabKota)
		wilayah[i].Provinsi = strings.TrimSpace(wilayah[i].Provinsi)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"success": true,
		"data": wilayah,
		"message": "Sukses",
		"param": gin.H{
			"query": nil,
		}})
}

func (h *ReferenceHandler) GetKodeWilPusat(c *gin.Context) {
	wilayah, err := h.service.GetKodeWil()
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "success": false, "error": err.Error()})
		return
	}

	filteredWilayah := []dto.ResponseKodeWil{}

	for _, w := range wilayah {
		w.KdWil = strings.TrimSpace(w.KdWil)
		w.NmWil = strings.TrimSpace(w.NmWil)
		w.AlWil = strings.TrimSpace(w.AlWil)
		w.KabKota = strings.TrimSpace(w.KabKota)
		w.Provinsi = strings.TrimSpace(w.Provinsi)

		if strings.HasSuffix(w.KdWil, "0") {
			filteredWilayah = append(filteredWilayah, w)
		}
	}

	c.JSON(200, gin.H{
		"code": 200,
		"success": true,
		"data": filteredWilayah,
		"message": "Sukses",
		"param": gin.H{
			"query": nil,
		},
	})
}

func (h *ReferenceHandler) GetKdJenisAB(c *gin.Context) {
	jenisAB, err := h.service.GetKdJenisAB()
	if err != nil {
		c.JSON(500, gin.H{"Code": 500, "Success" : false, "Message" : err.Error()})
		return
	}

	for i := range jenisAB {
		jenisAB[i].KdJenisAB = strings.TrimSpace(jenisAB[i].KdJenisAB)
		jenisAB[i].NmJenisAB = strings.TrimSpace(jenisAB[i].NmJenisAB)
	}

	c.JSON(200, gin.H{
		"code": "0000",
		"success": true,
		"message": "success",
		"data": jenisAB,
		"param": gin.H{
			"query": nil,
		},
	})
}


func (h *ReferenceHandler) GetKdMerek(c *gin.Context) {
	var req dto.RequestKdMerek
	err := c.ShouldBindQuery(&req); 
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
	if req.KdJenisAB == "" {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: "invalid param",
			Data: []string{},
			Param: req,
		}))
		return
	}

	listMerek, err := h.service.GetKdMerek(req.KdJenisAB)
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

	if listMerek == nil {
		c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
			Code: http.StatusNotFound,
			Success: false,
			Message: "kd jenis ab not found",
			Data: []string{},
			Param: req,
		}))
		return
	}

	c.JSON(200, gin.H{
		"code":"0000",
		"success": true,
		"message": "sukses",
		"data": listMerek,
		"param": req,
	})
}

func (h *ReferenceHandler) GetKdModel(c *gin.Context) {
	var req dto.RequestKdModel
	err := c.ShouldBindQuery(&req);
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
			Data: []string{},
			Param: req,
		}))
		return
	}

	if req.KdMerekAB == "" {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: "invalid param",
			Data: []string{},
			Param: req,
		}))
		return
	}

	listModel, err := h.service.GetKdModel(req.KdMerekAB)
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

	if listModel == nil {
		c.JSON(http.StatusNotFound, helper.Response(dto.ResponseParams{
			Code: http.StatusNotFound,
			Success: false,
			Message: "not found",
			Data: []string{},
			Param: req,
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "sukses",
		Data: listModel,
		Param: req,
	}))

}

func (h *ReferenceHandler) GetNilaiJual(c *gin.Context) {
	var req dto.RequestNilaiJual
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
			Data: []string{},
			Param: req,
		}))
		return
	}
	if req.KdMerekAB == "" {
		c.JSON(http.StatusBadRequest, helper.Response(dto.ResponseParams{
			Code: http.StatusBadRequest,
			Success: false,
			Message: "invalid param",
			Data: []string{},
			Param: req,
		}))
		return
	}

	listNilaiJual, err := h.service.GetNilaiJual(req.KdMerekAB)
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

	if listNilaiJual == nil {
		c.JSON(http.StatusNotFound, helper.Response(dto.ResponseParams{
			Code: http.StatusNotFound,
			Success: false,
			Message: "not found",
			Data: []string{},
			Param: req,
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "sukses",
		Data: listNilaiJual,
		Param: req,
	}))
}

func (h *ReferenceHandler) GetJenisBbm(c *gin.Context) {
	listBbm, err := h.service.GetJenisBbm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			Code: http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "sukses",
		Data: listBbm,
	}))
}

func (h *ReferenceHandler) GetJenisKepemilikan(c *gin.Context) {
	listKepemilikan, err := h.service.GetJenisKepemilikan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			Code: http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}))
		return
	}

	var result []dto.ResponseJenisKepemilikanParse
	for _, i := range listKepemilikan {
		idInt, _ :=strconv.Atoi(strings.TrimSpace(i.KdKepemilikan))
		result = append(result, dto.ResponseJenisKepemilikanParse{
			KdKepemilikan: idInt,
			NmKepemilikan: i.NmKepemilikan,
		}) 
	}
	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "sukses",
		Data: result,
	}))
}


func (h *ReferenceHandler) GetKdMohon(c *gin.Context) {
	listKdMohon, err := h.service.GetKdMohon()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.Response(dto.ResponseParams{
			Code: http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		}))
		return
	}

	c.JSON(http.StatusOK, helper.Response(dto.ResponseParams{
		Code: http.StatusOK,
		Success: true,
		Message: "sukses",
		Data: listKdMohon,
	}))
}
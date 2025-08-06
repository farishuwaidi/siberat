package helper

import (
	"Siberat/dto"
	"errors"
	"fmt"
	"log"
)

type ResponseWithData struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Param   any    `json:"param"` // Optional, if you want to include additional parameters
}

type ResponseWithoutData struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Response(params dto.ResponseParams) any {
	var response any
	var success bool
	var code string

	if params.Code >= 200 && params.Code < 300 {
		success = true
		code = "0000"
	} else {
		success = false
		code = toString(params.Code) // fallback ke kode asli jika bukan 2xx
	}

	if params.Data != nil {
		response = &ResponseWithData{
			Code:    code,
			Success: success,
			Data:    params.Data,
			Message: params.Message,
			Param:   params.Param,
		}
	} else {
		response = &ResponseWithoutData{
			Code:    code,
			Success: success,
			Message: params.Message,
		}
	}
	return response
}

func StringOfEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func toString(i int) string {
	return fmt.Sprintf("%d", i)
}

func ErrorResponseWithParam(code, message string, param any) any {
	return ResponseWithData{
		Code:    code,
		Success: false,
		Data: []any{},
		Message: message,
		Param: param,
	}
}

func ErrorValidation(message string) *ResponseWithData {
	return &ResponseWithData{
		Code:    "400",
		Success: false,
		Message: message,
		Data:    nil,
		Param:   nil,
	}
}

// Untuk error internal server (misal DB error, panic)
func ErrorInternal(message string, err error) *ResponseWithData {
	log.Printf("[INTERNAL ERROR] %s: %v", message, err)
	return &ResponseWithData{
		Code:    "500",
		Success: false,
		Message: message,
		Data:    nil,
		Param:   nil,
	}
}

// Untuk error jika data tidak ditemukan
func ErrorNotFound(message string) *ResponseWithData {
	return &ResponseWithData{
		Code:    "404",
		Success: false,
		Message: message,
		Data:    nil,
		Param:   nil,
	}
}

// Untuk response sukses dengan data
func Success(message string) *ResponseWithData {
	return &ResponseWithData{
		Code:    "200",
		Success: true,
		Message: message,
		Data:    nil,
		Param:   nil,
	}
}

// Untuk response sukses dengan data dan param tambahan
func SuccessWithData(message string, data any, param any) *ResponseWithData {
	return &ResponseWithData{
		Code:    "200",
		Success: true,
		Message: message,
		Data:    data,
		Param:   param,
	}
}

func NewCustomError(message string) error {
	return errors.New(message)
}

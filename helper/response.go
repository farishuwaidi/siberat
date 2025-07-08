package helper

import (
	"Siberat/dto"
	"fmt"
)

type ResponseWithData struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
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

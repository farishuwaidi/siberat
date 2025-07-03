package helper

import "Siberat/dto"

type ResponseWithData struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
	Param   any    `json:"param"` // Optional, if you want to include additional paraeters
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Response(params dto.ResponseParams) any {
	var response any
	var success bool

	if params.Code >= 200 && params.Code < 300 {
		success = true
	} else {
		success = false
	}

	if params.Data != nil {
		response = &ResponseWithData{
			Code:    params.Code,
			Success: success,
			Data:    params.Data,
			Message: params.Message, // Optional, if you want to include additional parameters
		}
	} else {
		response = &ResponseWithoutData{
			Code:    params.Code,
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
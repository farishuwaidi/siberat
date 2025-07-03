package dto

type ResponseParams struct {
	Code    int
	Success bool
	Data    any
	Message string
	Param   any
}
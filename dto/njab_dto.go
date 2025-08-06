package dto

type UpdateNjabRequest struct {
	IdUser   string                 `json:"id_user" binding:"required"`
	Method   string                 `json:"method" binding:"required"`
	DataSent map[string]interface{} `json:"data_sent" binding:"required"`
}

type GetNjabDetailsRequest struct {
	KdMerekAb string `json:"kd_merek_ab" form:"kd_merek_ab" binding:"required"`
}

type GetNilaiJualRequest struct {
	KdMerekAb string `json:"kd_merek_ab" form:"kd_merek_ab" binding:"required"`
	ThBuatan  string `json:"th_buatan" form:"th_buatan" binding:"required"`
}

type CrudNjabResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
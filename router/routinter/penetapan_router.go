package routinter

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func PenetapanRoute(api *gin.RouterGroup) {
	penetapanRepo := repository.NewPenetapanRepository(config.DB)
	penetapanService := service.NewPenetapanService(penetapanRepo)
	penetapanHandler := handler.NewPenetapanHandler(penetapanService)

	api.POST("/hitung-ulang", penetapanHandler.HitungUlang)
	api.POST("/get-penetapan", penetapanHandler.GetPenetapan)
}
package routinter

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func PendaftaranRoute(api *gin.RouterGroup) {
	pendaftaranRepo := repository.NewPendaftaranRepository(config.DB)
	pendaftaranService := service.NewPendaftaranService(pendaftaranRepo)
	pendaftaranHandler := handler.NewPendaftaranHandler(pendaftaranService)

	api.GET("/cek-transaksi", pendaftaranHandler.CekTransaksi)


}
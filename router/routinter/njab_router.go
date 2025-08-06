package routinter

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func NjabRouter(api *gin.RouterGroup) {
	njabRepo := repository.NewNjabRepository(config.DB)
	userRepo := repository.NewUserRepository(config.DB)
	njabService := service.NewNjabService(njabRepo, userRepo)
	njabHandler := handler.NewNjabHandler(njabService)

	api.POST("/crud-merek", njabHandler.UpdateMerek)
	api.POST("/get-detail-merek", njabHandler.GetMerekDetails)
	api.POST("/crud-nilai-jual", njabHandler.UpdateNilaiJual)
	api.POST("/get-detail-njab", njabHandler.GetDetailNjab)
}
package routinter

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func SubjekRoute(api *gin.RouterGroup) {
	subjekRepo := repository.NewSubjekRepository(config.DB)
	subjekService := service.NewSubjekService(subjekRepo)
	subjekHandler := handler.NewSubjekHandler(subjekService)

	api.GET("/get-subjek-ab", subjekHandler.GetSubjekPajak)
	api.GET("/get-kepemilikan-ab", subjekHandler.GetKepemilikanAb)
}
package router

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func ReferenceRoute(api *gin.RouterGroup) {
	referenceRepo := repository.NewReferenceRepository(config.DB)
	referenceService := service.NewReferenceService(referenceRepo)
	referenceHandler := handler.NewReferenceHandler(referenceService)

	api.GET("/reference/get-role", referenceHandler.GetKodeRole)
	api.GET("/reference/get-wiluppd", referenceHandler.GetKodeWil)
	api.GET("/reference/get-wiluppd-pusat", referenceHandler.GetKodeWilPusat)
}
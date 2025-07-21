package routinter

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func ObjekRoute(api *gin.RouterGroup) {
	objekRepo := repository.NewObjekRepository(config.DB)
	objekService := service.NewObjeKService(objekRepo)
	objekHandler := handler.NewObjekHandler(objekService)

	api.GET("/get-objek-by-nab", objekHandler.GetObjekByNab)
}
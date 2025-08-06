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
	penetapanRepo := repository.NewPenetapanRepository(config.DB)
	penetapanService := service.NewPenetapanService(penetapanRepo)
	objekService := service.NewObjeKService(objekRepo, penetapanService)
	objekHandler := handler.NewObjekHandler(objekService)

	api.GET("/get-objek-by-nab", objekHandler.GetObjekByNab)
	api.GET("/get-objek-subjek-by-nab", objekHandler.GetObjekSubjekByNab)
	api.GET("/get-info-pab-v2", objekHandler.GetInfoPAB)
}
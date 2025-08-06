package routinter

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
	api.GET("/reference/get-kdjenisab", referenceHandler.GetKdJenisAB)
	api.GET("/reference/get-kdmerekab", referenceHandler.GetKdMerek)
	api.GET("/reference/get-kdmodel", referenceHandler.GetKdModel)
	api.GET("/reference/get-nilai-jual", referenceHandler.GetNilaiJual)
	api.GET("/reference/get-jenis-bbm", referenceHandler.GetJenisBbm)
	api.GET("/reference/get-jenis-kepemilikan", referenceHandler.GetJenisKepemilikan)
	api.GET("/reference/get-kd-mohon", referenceHandler.GetKdMohon)
}
package router

import (
	"Siberat/router/routinter"

	"github.com/gin-gonic/gin"
)

func InitRouter(api *gin.RouterGroup) {
	routinter.AuthRouter(api)
	routinter.ObjekRoute(api)
	routinter.PendaftaranRoute(api)
	routinter.PenetapanRoute(api)
	routinter.ReferenceRoute(api)
	routinter.UserRouter(api)
}
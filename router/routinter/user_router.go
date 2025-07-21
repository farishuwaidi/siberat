package routinter

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(config.DB)
	UserService := service.NewUserService(userRepo)
	UserHandler := handler.NewUserHandler(UserService)

	// api.Use(middleware.JWTMiddleWare(), middleware.AuthorizeRoles("Petugas Bidang PP", "Petugas Bidang PSIP"))
	// api.Use(middleware.JWTMiddleWare())
	api.POST("/users", UserHandler.GetUserData)
	api.POST("/detail-users", UserHandler.GetUserDetail)
	api.POST("/update-users", UserHandler.UpdateUserData)
}
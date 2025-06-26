package router

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/middleware"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup) {
	userRepo := repository.NewUserRepository(config.DB)
	UserService := service.NewUserService(userRepo)
	UserHandler := handler.NewUserHandler(UserService)

	api.Use(middleware.JWTMiddleWare(), middleware.AuthorizeRoles("Petugas Bidang PP", "Petugas Bidang PSIP"))
	api.GET("/users", UserHandler.GetAllUser)
	api.PUT("/users/:id", UserHandler.UpdatedUser)
	api.GET("/users/:id", UserHandler.GetUserByID)
	api.DELETE("users/:id",UserHandler.DeleteUser)
}
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

	api.Use(middleware.JWTMiddleWare())
	api.GET("/users", UserHandler.GetAllUser)
}
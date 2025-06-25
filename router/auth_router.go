package router

import (
	"Siberat/config"
	"Siberat/handler"
	"Siberat/middleware"
	"Siberat/repository"
	"Siberat/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepo := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/login", authHandler.Login)
	api.Use(middleware.JWTMiddleWare())
	api.POST("/register",authHandler.Register)
}

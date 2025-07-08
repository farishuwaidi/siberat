package main

import (
	"Siberat/config"
	"Siberat/entity"
	"Siberat/router"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	config.DB.AutoMigrate(&entity.User{}, &entity.RoleUser{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // asal yang diizinkan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.AuthRouter(api)
	router.TestRoute(api)
	router.UserRouter(api)
	router.ReferenceRoute(api)

	r.Run(fmt.Sprintf("0.0.0.0:%v", config.ENV.PORT))
}
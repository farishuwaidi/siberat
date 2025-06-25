package main

import (
	"Siberat/config"
	"Siberat/entity"
	"Siberat/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	config.DB.AutoMigrate(&entity.User{}, &entity.RoleUser{})
	config.SeedRoles()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.AuthRouter(api)
	router.TestRoute(api)
	router.UserRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT)) // listen and serve on 0.0.0.0:8080
}
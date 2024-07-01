package cmd

import (
	"log"
	"os"

	"github.com/JerryJeager/weather-location-api/internal"
	middleware "github.com/JerryJeager/weather-location-api/middlware"
	"github.com/gin-gonic/gin"
)

func ExecuteApiRoutes() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.GET("/api/hello", internal.SayHello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}

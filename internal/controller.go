package internal

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	visitorName := ctx.Query("visitor_name")

	if visitorName == "" {
		ctx.JSON(http.StatusOK, gin.H{"message": "please provide 'visitor_name' with value as a query parameter"})
		return
	}

	ip := GetLocalIP()
	token := os.Getenv("PINFO_ACCESS_TOKEN")
	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	info, err := GetGeolocation(ip.String(), token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	weatherData, err := GetWeatherData(info.City, weatherApiKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"ip": ip.String(), "weather_info": *weatherData})

}

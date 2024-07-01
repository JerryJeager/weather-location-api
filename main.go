package main

import (
	"fmt"

	"github.com/JerryJeager/weather-location-api/cmd"
	"github.com/JerryJeager/weather-location-api/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	fmt.Println("starting the weather-location-server")
	cmd.ExecuteApiRoutes()
}

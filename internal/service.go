package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func GetGeolocation(ip string, apiKey string) (*IPinfoResponse, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location IPinfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}
	return &location, nil
}

func GetWeatherData(cityQuery string, apiKey string) (*WeatherResponse, error) {
	if cityQuery == "" {
		cityQuery = "auto"
	}
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json/key=%s&q=%s", apiKey, cityQuery)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}
	return &weatherData, nil
}

// func GenerateGreeting(name, location string, tempearture float32) string{

// }

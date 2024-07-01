package internal

type Response struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

type IPinfoResponse struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
}

type Location struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
}

type Current struct {
	TempC float32 `json:"temp_c"`
	TemF  float32 `json:"temp_f"`
}

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

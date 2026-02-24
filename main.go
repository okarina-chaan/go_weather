package main

import (
	"fmt"
	"net/http"
)

const (
	TokyoLat = 35.6895
	TokyoLon = 139.6917
)

func main() {
	latitude := TokyoLat
	longitude := TokyoLon
	weatherAPIURL := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m",
		latitude, longitude,
	)
	resp, err := http.Get(weatherAPIURL)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}

type WeatherResponse struct {
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Hourly    HourlyData `json:"hourly"`
}

type HourlyData struct {
	Time          []string  `json:"time"`
	Temperature2m []float64 `json:"temperature_2m"`
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// JSONをパースして構造体に格納
	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// 結果を表示
	fmt.Printf("Tokyo Weather:\n")
	fmt.Printf("Latitude: %f, Longitude: %f\n", weather.Latitude, weather.Longitude)
	fmt.Printf("First temperature: %f°C at %s\n",
		weather.Hourly.Temperature2m[0],
		weather.Hourly.Time[0])
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

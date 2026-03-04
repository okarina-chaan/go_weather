package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type City struct {
	Name      string
	Latitude  float64
	Longitude float64
}

var cities = []City{
	{Name: "Tokyo", Latitude: 35.6895, Longitude: 139.6917},
	{Name: "Osaka", Latitude: 34.6937, Longitude: 135.5023},
	{Name: "Sapporo", Latitude: 43.0642, Longitude: 141.3469},
	{Name: "Fukuoka", Latitude: 33.5904, Longitude: 130.4017},
	{Name: "New York", Latitude: 40.7128, Longitude: -74.0060},
	{Name: "London", Latitude: 51.5074, Longitude: -0.1278},
	{Name: "Paris", Latitude: 48.8566, Longitude: 2.3522},
	{Name: "Sydney", Latitude: -33.8688, Longitude: 151.2093},
	{Name: "Singapore", Latitude: 1.3521, Longitude: 103.8198},
	{Name: "Dubai", Latitude: 25.2048, Longitude: 55.2708},
}

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

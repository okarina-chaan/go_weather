package main

import (
	"fmt"
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
	fmt.Println(weatherAPIURL)
}

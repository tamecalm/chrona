package tz

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherInfo struct {
	Temperature float64
	Condition   string
	Sunrise     string
	Sunset      string
	IsDay       bool
}

var weatherCodes = map[int]string{
	0: "Clear Sky", 1: "Mainly Clear", 2: "Partly Cloudy", 3: "Overcast",
	45: "Fog", 48: "Rime Fog", 51: "Light Drizzle", 53: "Moderate Drizzle",
	55: "Dense Drizzle", 61: "Light Rain", 63: "Moderate Rain", 65: "Heavy Rain",
	71: "Light Snow", 73: "Moderate Snow", 75: "Heavy Snow",
	80: "Rain Showers", 81: "Moderate Showers", 82: "Violent Showers",
	95: "Thunderstorm", 96: "Thunderstorm + Hail",
}

func FetchWeather(lat, lon float64) (*WeatherInfo, error) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true&daily=sunrise,sunset&timezone=auto",
		lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	current := data["current_weather"].(map[string]any)
	daily := data["daily"].(map[string]any)
	code := int(current["weathercode"].(float64))

	return &WeatherInfo{
		Temperature: current["temperature"].(float64),
		Condition:   weatherCodes[code],
		Sunrise:     daily["sunrise"].([]any)[0].(string),
		Sunset:      daily["sunset"].([]any)[0].(string),
		IsDay:       current["is_day"].(float64) == 1,
	}, nil
}

func fetchAPITime(timezone string) (string, error) {
	url := fmt.Sprintf("https://worldtimeapi.org/api/timezone/%s", timezone)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	datetime, ok := data["datetime"].(string)
	if !ok {
		return "", fmt.Errorf("invalid datetime format")
	}

	return datetime, nil
}

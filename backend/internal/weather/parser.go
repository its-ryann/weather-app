package weather

import "github.com/its-ryann/weather-app/internal/models"

// RawWeatherData represents the shape of data coming from the external weather API.
type RawWeatherData struct {
	CityName    string
	TempCelsius float64
	Condition   string
	Humidity    int
}

// ParseWeather converts raw external API data into our internal WeatherResponse model.
func ParseWeather(raw RawWeatherData) models.WeatherResponse {
	return models.WeatherResponse{
		City:      raw.CityName,
		TempC:     raw.TempCelsius,
		Condition: raw.Condition,
		Humidity:  raw.Humidity,
	}
}

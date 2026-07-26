package models

// WeatherResponse is our internal representation returned by the API.
type WeatherResponse struct {
	City      string  `json:"city"`
	TempC     float64 `json:"tempC"`
	Condition string  `json:"condition"`
	Humidity  int     `json:"humidity"`
}

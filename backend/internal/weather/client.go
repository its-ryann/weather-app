package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/its-ryann/weather-app/internal/models"
)

// Client talks to the external weather API (OpenWeatherMap-shaped responses).
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// owmResponse mirrors the subset of OpenWeatherMap's response shape we care about.
type owmResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// GetWeather fetches current weather for a city and returns our internal model.
func (c *Client) GetWeather(city string) (models.WeatherResponse, error) {
	url := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=metric", c.baseURL, city, c.apiKey)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return models.WeatherResponse{}, fmt.Errorf("calling weather API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return models.WeatherResponse{}, fmt.Errorf("city %q not found", city)
	}
	if resp.StatusCode != http.StatusOK {
		return models.WeatherResponse{}, fmt.Errorf("weather API returned status %d", resp.StatusCode)
	}

	var raw owmResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return models.WeatherResponse{}, fmt.Errorf("decoding weather response: %w", err)
	}

	condition := ""
	if len(raw.Weather) > 0 {
		condition = raw.Weather[0].Main
	}

	return ParseWeather(RawWeatherData{
		CityName:    raw.Name,
		TempCelsius: raw.Main.Temp,
		Condition:   condition,
		Humidity:    raw.Main.Humidity,
	}), nil
}

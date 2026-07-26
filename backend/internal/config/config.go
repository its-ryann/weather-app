package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherAPIKey     string
	WeatherAPIBaseURL string
	Port              string
}

func Load() (*Config, error) {
	// Ignore error: in production, real env vars are already set and no .env file exists.
	_ = godotenv.Load()

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY is required but not set")
	}

	baseURL := os.Getenv("WEATHER_API_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.openweathermap.org/data/2.5"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		WeatherAPIKey:     apiKey,
		WeatherAPIBaseURL: baseURL,
		Port:              port,
	}, nil
}

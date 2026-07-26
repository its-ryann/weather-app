package main

import (
	"log"
	"net/http"

	"github.com/its-ryann/weather-app/internal/config"
	"github.com/its-ryann/weather-app/internal/handler"
	"github.com/its-ryann/weather-app/internal/weather"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	weatherClient := weather.NewClient(cfg.WeatherAPIBaseURL, cfg.WeatherAPIKey)
	weatherHandler := handler.NewWeatherHandler(weatherClient)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/weather", weatherHandler.GetWeather)
	mux.HandleFunc("GET /healthz", handler.HealthCheck)

	log.Printf("server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

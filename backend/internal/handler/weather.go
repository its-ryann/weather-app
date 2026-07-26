package handler

import (
	"encoding/json"
	"net/http"

	"github.com/its-ryann/weather-app/internal/models"
)

// WeatherService is the interface the handler depends on — satisfied by both
// the real weather.Client and any test mock.
type WeatherService interface {
	GetWeather(city string) (models.WeatherResponse, error)
}

type WeatherHandler struct {
	service WeatherService
}

func NewWeatherHandler(service WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		writeJSONError(w, http.StatusBadRequest, "city parameter required")
		return
	}

	result, err := h.service.GetWeather(city)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

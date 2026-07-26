package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/its-ryann/weather-app/internal/models"
)

// mockWeatherService lets us test the handler without a real weather client.
type mockWeatherService struct {
	response models.WeatherResponse
	err      error
}

func (m *mockWeatherService) GetWeather(city string) (models.WeatherResponse, error) {
	return m.response, m.err
}

func TestWeatherHandler_Success(t *testing.T) {
	mock := &mockWeatherService{
		response: models.WeatherResponse{
			City: "Nairobi", TempC: 24.5, Condition: "Clear", Humidity: 61,
		},
	}
	h := NewWeatherHandler(mock)

	req := httptest.NewRequest(http.MethodGet, "/api/weather?city=Nairobi", nil)
	rec := httptest.NewRecorder()

	h.GetWeather(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

func TestWeatherHandler_MissingCityParam(t *testing.T) {
	mock := &mockWeatherService{}
	h := NewWeatherHandler(mock)

	req := httptest.NewRequest(http.MethodGet, "/api/weather", nil)
	rec := httptest.NewRecorder()

	h.GetWeather(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

func TestWeatherHandler_CityNotFound(t *testing.T) {
	mock := &mockWeatherService{err: errors.New("city not found")}
	h := NewWeatherHandler(mock)

	req := httptest.NewRequest(http.MethodGet, "/api/weather?city=Nowhereland", nil)
	rec := httptest.NewRecorder()

	h.GetWeather(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}

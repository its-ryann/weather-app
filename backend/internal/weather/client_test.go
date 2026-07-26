package weather

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetWeather_Success(t *testing.T) {
	// Fake server mimicking OpenWeatherMap's response shape
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"name": "Nairobi",
			"main": {"temp": 24.5, "humidity": 61},
			"weather": [{"main": "Clear"}]
		}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "fake-api-key")

	result, err := client.GetWeather("Nairobi")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.City != "Nairobi" {
		t.Errorf("expected city Nairobi, got %s", result.City)
	}
	if result.TempC != 24.5 {
		t.Errorf("expected tempC 24.5, got %f", result.TempC)
	}
	if result.Condition != "Clear" {
		t.Errorf("expected condition Clear, got %s", result.Condition)
	}
	if result.Humidity != 61 {
		t.Errorf("expected humidity 61, got %d", result.Humidity)
	}
}

func TestClient_GetWeather_CityNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "city not found"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "fake-api-key")

	_, err := client.GetWeather("Nowhereland")
	if err == nil {
		t.Fatal("expected an error for unknown city, got nil")
	}
}

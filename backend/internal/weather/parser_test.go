package weather

import "testing"

func TestParseWeather(t *testing.T) {
	raw := RawWeatherData{
		CityName:    "Nairobi",
		TempCelsius: 24.5,
		Condition:   "Clear",
		Humidity:    61,
	}

	result := ParseWeather(raw)

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

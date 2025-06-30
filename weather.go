package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (app *application) calcLocation(city string) (coOrdinate, error) {
	searchURL := fmt.Sprintf("%vq=%v&limit=5&appid=%v", app.BASE_URL, city, app.API_KEY)

	response, err := http.Get(searchURL)
	if err != nil {
		return coOrdinate{}, fmt.Errorf("GET request failed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return coOrdinate{}, fmt.Errorf("Status code: %v", response.StatusCode)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return coOrdinate{}, fmt.Errorf("IO.ReadAll failed: %w", err)
	}

	xy := coOrdinate{}

	err = json.Unmarshal(bytes, &xy)
	if err != nil {
		return coOrdinate{}, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return xy, nil
}

func (app *application) reportWeather(lat, lon float64) (weatherInfo, error) {
	searchURL := fmt.Sprintf("%vlat=%v&lon=%v&appid=%v", app.BASE_URL, lat, lon, app.API_KEY)

	response, err := http.Get(searchURL)
	if err != nil {
		return weatherInfo{}, fmt.Errorf("GET request failed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return weatherInfo{}, fmt.Errorf("Status code: %v", response.StatusCode)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return weatherInfo{}, fmt.Errorf("IO.ReadAll failed: %w", err)
	}

	weather := weatherInfo{}

	err = json.Unmarshal(bytes, &weather)
	if err != nil {
		return weatherInfo{}, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return weather, nil
}

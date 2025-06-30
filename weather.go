package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (app *application) reportWeather(city string) (weatherInfo, error) {
	searchURL, err := url.Parse(app.BASE_URL)
	if err != nil {
		return weatherInfo{}, fmt.Errorf("invalid base URL: %w", err)
	}
	params := url.Values{}
	params.Set("q", city)
	params.Set("appid", app.API_KEY)
	searchURL.RawQuery = params.Encode()

	response, err := http.Get(searchURL.String())
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

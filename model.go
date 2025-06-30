package main

type application struct {
	BASE_URL string
	API_KEY  string
}

type weatherInfo struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp          float64 `json:"temp"`
		TempFeelsLike float64 `json:"feels_like"`
		Humidity      int     `json:"humidity"`
	} `json:"main"`
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/lpernett/godotenv"
)

type coOrdinate struct {
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
}

type printWeather struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func calcLocation(BASE_URL, API_KEY, city string) (float64, float64) {
	searchURL := fmt.Sprintf("%vq=%v&limit=5&appid=%v", BASE_URL, city, API_KEY)

	response, err := http.Get(searchURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)

		xy := coOrdinate{}

		json.Unmarshal(bytes, &xy)

		return xy.Coord.Lat, xy.Coord.Lon
	} else {
		fmt.Println("Failed. Status:", response.Status)
		return -1, -1
	}
}

func reportWeather(BASE_URL, API_KEY string, lat, lon float64) {
	searchURL := fmt.Sprintf("%vlat=%v&lon=%v&appid=%v", BASE_URL, lat, lon, API_KEY)

	response, err := http.Get(searchURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)

		weather := printWeather{}

		json.Unmarshal(bytes, &weather)

		fmt.Println("Current:", weather.Weather[0].Main)
		fmt.Println("Description:", weather.Weather[0].Description)
		fmt.Printf("Temperature:%0.1f\n", weather.Main.Temp-273.15)
		fmt.Printf("Humidity:%v%%\n", weather.Main.Humidity)
	} else {
		fmt.Println("Failed. Status:", response.Status)
	}
}

func main() {
	loadEnv()

	BASE_URL := os.Getenv("BASE_URL")
	API_KEY := os.Getenv("API_KEY")

	fmt.Print("Enter Location: ")
	var city string
	fmt.Scanln(&city)

	lat, lon := calcLocation(BASE_URL, API_KEY, city)

	reportWeather(BASE_URL, API_KEY, lat, lon)
}

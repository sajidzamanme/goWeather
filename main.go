package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lpernett/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	loadEnv()

	app := &application{
		BASE_URL: os.Getenv("BASE_URL"),
		API_KEY:  os.Getenv("API_KEY"),
	}

	fmt.Print("Enter Location: ")

	reader := bufio.NewReader(os.Stdin)
	city, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	city = strings.TrimSpace(city)

	weather, err := app.reportWeather(city)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current:", weather.Weather[0].Main)
	fmt.Println("Description:", weather.Weather[0].Description)
	fmt.Printf("Temperature: %.1f°C\n", weather.Main.Temp-273.15)
	fmt.Printf("Feels Like: %.1f°C\n", weather.Main.TempFeelsLike-273.15)
	fmt.Printf("Humidity:%v%%\n", weather.Main.Humidity)
}

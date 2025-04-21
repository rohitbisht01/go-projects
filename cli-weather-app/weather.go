package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aquasecurity/table"
	"github.com/briandowns/spinner"
)

type FullWeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Main struct {
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
		SeaLevel int64   `json:"sea_level"`
	} `json:"main"`

	Sys struct {
		Country string `json:"country"`
	} `json:"sys"`

	Visibility int64  `json:"visibility"`
	Name       string `json:"name"`
}

type WeatherType struct {
	Coord struct {
		Lon float64
		Lat float64
	}
	Name               string
	Country            string
	WeatherDescription string
	Main               struct {
		Pressure float64
		Humidity float64
		SeaLevel int64
	}
	Visibility int64
}

func Fetch(city string) {
	WEATHER_API_KEY := os.Getenv("WEATHER_API_KEY")
	if WEATHER_API_KEY == "" {
		log.Fatal("Missing WEATHER_API_KEY environment variable")
	}

	baseUrl := "http://api.openweathermap.org/data/2.5/weather?" + "appid=" + WEATHER_API_KEY + "&q=" + city

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Fetching weather data..."
	s.Start()

	resp, err := http.Get(baseUrl)

	s.Stop()

	if err != nil {
		log.Fatal("Error while making HTTP request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error while reading response body:", err)
	}

	var fullResp FullWeatherResponse
	err = json.Unmarshal(data, &fullResp)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	var weather WeatherType
	weather.Coord.Lat = fullResp.Coord.Lat
	weather.Coord.Lon = fullResp.Coord.Lon
	weather.Name = fullResp.Name
	weather.Country = fullResp.Sys.Country
	weather.Visibility = fullResp.Visibility
	weather.Main.Pressure = fullResp.Main.Pressure
	weather.Main.Humidity = fullResp.Main.Humidity
	weather.Main.SeaLevel = fullResp.Main.SeaLevel
	if len(fullResp.Weather) > 0 {
		weather.WeatherDescription = fullResp.Weather[0].Description
	}

	Print(weather)
}

func Print(weather WeatherType) {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("City", "Weather Description", "Lat", "Lon", "Visibility", "Pressure", "Humidity", "SeaLevel")

	tbl.AddRow(
		weather.Name,
		weather.WeatherDescription,
		fmt.Sprintf("%.4f", weather.Coord.Lat),
		fmt.Sprintf("%.4f", weather.Coord.Lon),
		fmt.Sprintf("%d", weather.Visibility),
		fmt.Sprintf("%.4f", weather.Main.Pressure),
		fmt.Sprintf("%.4f", weather.Main.Humidity),
		fmt.Sprintf("%d", weather.Main.SeaLevel),
	)

	tbl.Render()
}

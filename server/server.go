package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/sanamlimbu/getweather/getweather"
	"google.golang.org/grpc"
)

const openMeteoAPI = "https://api.open-meteo.com/v1/forecast"

type weatherServer struct {
	getweather.UnimplementedWeatherServer
	httpClient *http.Client
}

func newWeatherServer() *weatherServer {
	s := &weatherServer{httpClient: &http.Client{}}
	return s
}

// Response from Open-Meteo
type openMeteoResponse struct {
	Latitude       float32 `json:"latitude"`
	Longitude      float32 `json:"longitude"`
	CurrentWeather struct {
		Temperature   float32 `json:"temperature"`
		Windspeed     float32 `json:"windspeed"`
		Winddirection float32 `json:"winddirection"`
		Time          string  `json:"time"`
	} `json:"current_weather"`
}

func (s *weatherServer) CurrentWeatherInfo(ctx context.Context, location *getweather.Location) (*getweather.WeatherInfo, error) {
	req, err := http.NewRequest("GET", openMeteoAPI, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("latitude", fmt.Sprint(location.Latitude))
	q.Add("longitude", fmt.Sprint(location.Longitude))
	q.Add("current_weather", fmt.Sprint(true))
	req.URL.RawQuery = q.Encode()

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode json response
	target := &openMeteoResponse{}
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return nil, err
	}

	weatherInfo := &getweather.WeatherInfo{
		Location: &getweather.Location{
			Latitude:  target.Latitude,
			Longitude: target.Longitude,
		},
		Temperature:   target.CurrentWeather.Temperature,
		Windspeed:     target.CurrentWeather.Windspeed,
		Winddirection: target.CurrentWeather.Winddirection,
		Time:          target.CurrentWeather.Time,
	}

	return weatherInfo, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listenig on port 9001 ......")
	grpcServer := grpc.NewServer()
	getweather.RegisterWeatherServer(grpcServer, newWeatherServer())
	grpcServer.Serve(listener)
}

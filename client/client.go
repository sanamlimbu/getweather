package main

import (
	"context"
	"log"
	"time"

	"github.com/sanamlimbu/getweather/getweather"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := getweather.NewWeatherClient(conn)

	locations := map[string]*getweather.Location{
		"Kathmandu": {
			Latitude:  27.70,
			Longitude: 85.32,
		},
		"Perth": {
			Latitude:  -31.95,
			Longitude: 115.86,
		},
		"New York": {
			Latitude:  40.71,
			Longitude: -74.01,
		},
	}

	for k, location := range locations {
		printCurrentWeather(client, k,
			&getweather.Location{
				Latitude:  location.Latitude,
				Longitude: location.Longitude,
			})
	}
}

// printCurrentWeather gets the current weatehr information of given location
func printCurrentWeather(client getweather.WeatherClient, city string, location *getweather.Location) {
	log.Printf("\nGetting current weather info of %s (%f, %f)", city, location.Latitude, location.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	weatherInfo, err := client.CurrentWeatherInfo(ctx, location)
	if err != nil {
		log.Fatalf("client.CurrentWeatherInfo failed: %v", err)
	}
	log.Println(weatherInfo)
}

package openweather

import (
	"net/http"
	"time"
)

type OpenWeatherApi struct {
	apiKey       string
	measureUnits string
	webClient    http.Client
}

func Initialize(apiKey string) *OpenWeatherApi {
	api := OpenWeatherApi{
		apiKey: apiKey,
		webClient: http.Client{
			Timeout: time.Minute,
		},
		measureUnits: STANDARD,
	}
	return &api
}

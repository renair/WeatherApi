package openweather

import (
	"net/http"
	"time"

	"github.com/renair/weather/persistence"
)

type OpenWeatherApi struct {
	apiKey       string
	measureUnits string
	webClient    http.Client
	storage      *persistence.Storage
}

const CACHE_TIME = 30 * 60

func Initialize(apiKey string) *OpenWeatherApi {
	api := OpenWeatherApi{
		apiKey: apiKey,
		webClient: http.Client{
			Timeout: time.Minute,
		},
		measureUnits: STANDARD,
		storage:      nil,
	}
	return &api
}

func (owa *OpenWeatherApi) SetStorage(s *persistence.Storage) {
	owa.storage = s
}

func (owa *OpenWeatherApi) SetMeasureUnits(units string) {
	owa.measureUnits = units
}

package resolver

import (
	"time"

	"github.com/renair/weather/models"
	"github.com/renair/weather/openweather"
)

func convertWeatherFromApi(w openweather.WeatherData, loc models.Location) *models.WeatherData {
	res := models.WeatherData{
		Location: loc,
		Values: models.MainValues{
			Temperature: w.GeneralMetrics.Temperature,
			Pressure:    &w.GeneralMetrics.Pressure,
			Humidity:    w.GeneralMetrics.Humidity,
		},
		Cloud: models.Cloud{
			Cloudiness: 123,
			IsRain:     toBoolPtr(w.Rain.OneHour > 0),
			IsSnow:     toBoolPtr(w.Snow.OneHour > 0),
		},
		Wind: nil,
		Date: time.Unix(w.MeasuredAt, 0),
	}
	return &res
}

func toBoolPtr(val bool) *bool {
	return &val
}

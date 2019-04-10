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
			IsRain: toBoolPtr(w.Rain.OneHour > 0),
			IsSnow: toBoolPtr(w.Snow.OneHour > 0),
		},
		Wind: nil,
		Date: time.Unix(w.MeasuredAt, 0),
	}
	return &res
}

func convertForecastFromApi(w openweather.ForecastData, loc models.Location) []models.WeatherData {
	res := make([]models.WeatherData, w.Count)
	for i, weather := range w.Measures {
		res[i] = models.WeatherData{
			Location: loc,
			Values: models.MainValues{
				Temperature: weather.GeneralMetrics.Temperature,
				Pressure:    &weather.GeneralMetrics.Pressure,
				Humidity:    weather.GeneralMetrics.Humidity,
			},
			Cloud: models.Cloud{
				IsRain: toBoolPtr(weather.Rain.OneHour > 0),
				IsSnow: toBoolPtr(weather.Snow.OneHour > 0),
			},
			Wind: nil,
			Date: time.Unix(weather.MeasuredAt, 0),
		}
	}
	return res
}

func toBoolPtr(val bool) *bool {
	return &val
}

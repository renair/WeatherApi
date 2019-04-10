package resolver

import (
	"context"
	"fmt"

	"github.com/renair/weather"
	"github.com/renair/weather/models"

	"github.com/renair/weather/openweather"
	"github.com/renair/weather/persistence"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	ApiClient   *openweather.OpenWeatherApi
	Persistance *persistence.Storage
}

func (r *Resolver) Mutation() weather.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() weather.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) WeatherData() weather.WeatherDataResolver {
	return &weatherDataResolwer{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateLocation(ctx context.Context, input models.NewLocation) (*models.Location, error) {
	return r.Persistance.SaveNewLocation(input)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Location(ctx context.Context, id int) (*models.Location, error) {
	return r.Persistance.LocationById(id)
}

func (r *queryResolver) LocationsInRegion(ctx context.Context, longitude float64, latitude float64, radius float64) ([]models.Location, error) {
	return r.Persistance.LocationsInRadius(longitude, latitude, radius)
}

func (r *queryResolver) WeatherInRegion(ctx context.Context, longitude float64, latitude float64, radius float64) ([]models.WeatherData, error) {
	locations, err := r.Persistance.LocationsInRadius(longitude, latitude, radius)
	if err != nil {
		return nil, err
	}
	res := make([]models.WeatherData, len(locations))
	for i, loc := range locations {
		weatherData, err := r.WeatherInLocation(ctx, loc.ID)
		if err != nil {
			continue
		}
		res[i] = *weatherData
	}
	return res, nil
}

func (r *queryResolver) WeatherInLocation(ctx context.Context, locationID int) (*models.WeatherData, error) {
	loc, err := r.Persistance.LocationById(locationID)
	if err != nil {
		return nil, err
	}
	weather, err := r.ApiClient.GetCurrentWeatherByCoords(float32(loc.Longitude), float32(loc.Latitude))
	return convertWeatherFromApi(weather, *loc), err
}

type weatherDataResolwer struct{ *Resolver }

func (r *weatherDataResolwer) Forecast(ctx context.Context, obj *models.WeatherData) ([]models.WeatherData, error) {
	forecast, err := r.ApiClient.GetForecastByCoords(float32(obj.Location.Longitude), float32(obj.Location.Latitude))
	if err != nil {
		fmt.Println(err.Error())
		return []models.WeatherData{}, err
	}
	fmt.Println(forecast)
	return convertForecastFromApi(forecast, obj.Location), nil
}

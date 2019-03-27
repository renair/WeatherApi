package resolver

import (
	"context"
	"fmt"

	"github.com/renair/weather"
	"github.com/renair/weather/models"
	"github.com/renair/weather/openweather"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	nextLocationId int
	locations      []models.Location
	ApiClient      *openweather.OpenWeatherApi
}

func (r *Resolver) Mutation() weather.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() weather.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateLocation(ctx context.Context, input models.NewLocation) (*models.Location, error) {
	r.nextLocationId += 1
	loc := models.Location{
		Longitude:    input.Longitude,
		Latitude:     input.Latitude,
		LocationName: input.LocationName,
		ID:           r.nextLocationId,
	}
	r.locations = append(r.locations, loc)
	return &loc, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Location(ctx context.Context, id int) (*models.Location, error) {
	index := id - 1
	if index >= len(r.locations) {
		return nil, fmt.Errorf("There is no location with this id %d", id)
	}
	return &r.locations[index], nil
}
func (r *queryResolver) LocationsInRegion(ctx context.Context, longitude *float64, latitude *float64, radius *float64) ([]models.Location, error) {
	panic("not implemented")
}
func (r *queryResolver) WeatherInRegion(ctx context.Context, longitude *float64, latitude *float64, radius *float64) ([]*models.WeatherData, error) {
	panic("not implemented")
}
func (r *queryResolver) WeatherInLocation(ctx context.Context, locationID int) (*models.WeatherData, error) {
	index := locationID - 1
	if index >= len(r.locations) {
		return nil, fmt.Errorf("There is no location with this id %d", locationID)
	}
	loc := r.locations[index]
	weather, err := r.ApiClient.GetCurrentWeatherByCoords(float32(loc.Longitude), float32(loc.Latitude))
	return convertWeatherFromApi(weather, loc), err
}

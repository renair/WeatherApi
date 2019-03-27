package resolver

import (
	"context"

	"github.com/renair/weather"
	"github.com/renair/weather/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	locations []models.Location
}

func (r *Resolver) Mutation() weather.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() weather.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateLocation(ctx context.Context, input models.NewLocation) (*models.Location, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Location(ctx context.Context, id int) (*models.Location, error) {
	panic("not implemented")
}
func (r *queryResolver) LocationsInRegion(ctx context.Context, longitude *float64, latitude *float64, radius *float64) ([]models.Location, error) {
	panic("not implemented")
}
func (r *queryResolver) WeatherInRegion(ctx context.Context, longitude *float64, latitude *float64, radius *float64) ([]*models.WeatherData, error) {
	panic("not implemented")
}
func (r *queryResolver) WeatherInLocation(ctx context.Context, locationID int) (*models.WeatherData, error) {
	panic("not implemented")
}

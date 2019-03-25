package weather

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateLocation(ctx context.Context, input NewLocation) (*Location, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Location(ctx context.Context, id int) (*Location, error) {
	panic("not implemented")
}
func (r *queryResolver) LocationsInRegion(ctx context.Context, longitude *float64, latitude *float64, radius *float64) ([]Location, error) {
	panic("not implemented")
}
func (r *queryResolver) WeatherInRegion(ctx context.Context, longitude *float64, latitude *float64, radius *float64) ([]*WeatherData, error) {
	panic("not implemented")
}
func (r *queryResolver) WeatherInLocation(ctx context.Context, locationID int) (*WeatherData, error) {
	panic("not implemented")
}

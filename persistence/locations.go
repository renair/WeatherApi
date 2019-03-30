package persistence

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/renair/weather/models"
)

func (s *Storage) SaveNewLocation(loc *models.Location) error {
	newId, err := s.conn.Incr(idKey).Result()
	if err != nil {
		return fmt.Errorf("Can't create new location. Error in obtaining key: %s", err.Error())
	}

	stringNewId := fmt.Sprint(newId)
	changed, err := s.conn.GeoAdd(locationsCoordsStorage, &redis.GeoLocation{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		Name:      stringNewId,
	}).Result()
	if err != nil || changed == 0 {
		return fmt.Errorf("Error when storing new location: %v", err)
	}

	if loc.LocationName != nil {
		isStored, err := s.conn.HSet(locationNamesStorage, stringNewId, *loc.LocationName).Result()
		if err != nil || !isStored {
			return fmt.Errorf("Error when storing location name: %v", err)
		}
	}

	loc.ID = int(newId)
	return nil
}

func (s *Storage) LocationById(id int) (*models.Location, error) {
	stringId := fmt.Sprint(id)
	positions, err := s.conn.GeoPos(locationsCoordsStorage, stringId).Result()
	if err != nil || len(positions) != 1 {
		return nil, err
	}
	locationName, err := s.conn.HGet(locationNamesStorage, stringId).Result()
	if err != nil {
		return nil, err
	}
	return &models.Location{
		LocationName: &locationName,
		Longitude:    positions[0].Longitude,
		Latitude:     positions[0].Latitude,
		ID:           id,
	}, nil
}

func (s *Storage) LocationsInRadius(lon float64, lat float64, radius float64) ([]*models.Location, error) {
	savedLocs, err := s.conn.GeoRadius(locationsCoordsStorage, lon, lat, &redis.GeoRadiusQuery{
		Radius: radius,
		Unit:   "km",
	}).Result()
	if err != nil {
		return nil, err
	}

	res := make([]*models.Location, len(savedLocs))
	for i, loc := range savedLocs {
		locationName, _ := s.conn.HGet(locationNamesStorage, loc.Name).Result()
		locId, _ := strconv.ParseInt(loc.Name, 10, 64)
		res[i] = &models.Location{
			LocationName: &locationName,
			Longitude:    loc.Longitude,
			Latitude:     loc.Latitude,
			ID:           int(locId),
		}
	}

	return res, nil
}

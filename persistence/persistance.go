package persistence

import (
	"fmt"

	"github.com/go-redis/redis"
)

const (
	idKey                  = "location_id"
	locationsCoordsStorage = "locationsCoords"
	locationNamesStorage   = "locationsNames"
)

type Storage struct {
	conn *redis.Client
}

func NewStorage(adress string, port string) *Storage {
	s := Storage{
		conn: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", adress, port),
			DB:   0,
		}),
	}
	return &s
}

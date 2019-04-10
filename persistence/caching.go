package persistence

import (
	"fmt"
	"time"
)

func (s *Storage) CacheFor(key string, value string, second int) bool {
	res, err := s.conn.SetNX(key, value, time.Duration(second)*time.Second).Result()
	if err != nil {
		fmt.Println("Error in caching: ", err.Error())
	}
	return res
}

func (s *Storage) GetCached(key string) (string, bool) {
	res, err := s.conn.Get(key).Result()
	if err != nil {
		return "", false
	}
	return res, true
}

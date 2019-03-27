package models

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type WeatherData struct {
	Location Location   `json:"location"`
	Values   MainValues `json:"values"`
	Cloud    Cloud      `json:"cloud"`
	Wind     *Wind      `json:"wind"`
	Date     time.Time  `json:"date"`
}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix()
	if timestamp < 0 {
		timestamp = 0
	}
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}

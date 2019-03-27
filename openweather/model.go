package openweather

type WeatherData struct {
	Coords         CoordsData        `json:"coord"`
	GeneralData    []ReadableWeather `json:"weather"`
	Base           string            `json:"base"`
	GeneralMetrics EnvironmentalData `json:"main"`
	MeasuredAt     int64             `json:"dt"`
	Rain           Precipitation     `json:"rain"`
	Snow           Precipitation     `json:"snow"`
	CityId         int               `json:"id"`
}

type CoordsData struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type ReadableWeather struct {
	Id          int    `json:"id"`
	Name        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type EnvironmentalData struct {
	Temperature         float64 `json:"temp"`
	Pressure            float64 `json:"pressure"`
	Humidity            int     `json:"humidity"`
	MinTemperature      float64 `json:"temp_min"`
	MaxTemperature      float64 `json:"temp_max"`
	SeaLevelPressure    float64 `json:"sea_level"`
	GroundLevelPressure float64 `json:"grnd_level"`
}

type WindData struct {
	Speed     float64 `json:"speed"`
	Direction float64 `json:"deg"`
}

type Precipitation struct {
	OneHour   float64 `json:"1h"`
	ThreeHour float64 `json:"3h"`
}

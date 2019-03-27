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
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}

type ReadableWeather struct {
	Id          int    `json:"id"`
	Name        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type EnvironmentalData struct {
	Temperature         float32 `json:"temp"`
	Pressure            float32 `json:"pressure"`
	Humidity            int     `json:"humidity"`
	MinTemperature      float32 `json:"temp_min"`
	MaxTemperature      float32 `json:"temp_max"`
	SeaLevelPressure    float32 `json:"sea_level"`
	GroundLevelPressure float32 `json:"grnd_level"`
}

type WindData struct {
	Speed     float32 `json:"speed"`
	Direction float32 `json:"deg"`
}

type Precipitation struct {
	OneHour   int `json:"1h"`
	ThreeHour int `json:"3h"`
}

package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (owa *OpenWeatherApi) GetCurrentWeatherByCoords(lon float32, lat float32) (WeatherData, error) {
	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather", nil)
	if err != nil {
		return WeatherData{}, fmt.Errorf("Can't create http request")
	}

	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%.2f", lat))
	q.Add("lon", fmt.Sprintf("%.2f", lon))
	q.Add("appid", owa.apiKey)
	q.Add("units", owa.measureUnits)
	req.URL.RawQuery = q.Encode()

	resp, err := owa.webClient.Do(req)
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()
	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherData{}, err
	}

	var result WeatherData
	err = json.Unmarshal(resData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (owa *OpenWeatherApi) GetForecastByCoords(lon float32, lat float32) (ForecastData, error) {
	req, err := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/forecast", nil)
	if err != nil {
		return ForecastData{}, fmt.Errorf("Can't create http request")
	}

	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%.2f", lat))
	q.Add("lon", fmt.Sprintf("%.2f", lon))
	q.Add("appid", owa.apiKey)
	q.Add("units", owa.measureUnits)
	req.URL.RawQuery = q.Encode()

	resp, err := owa.webClient.Do(req)
	if err != nil {
		return ForecastData{}, err
	}
	defer resp.Body.Close()
	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ForecastData{}, err
	}

	var result ForecastData
	err = json.Unmarshal(resData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

package openweather

import (
	"fmt"

	"io/ioutil"
	"net/http"
)

func (owa *OpenWeatherApi) makeRequest(req *http.Request) ([]byte, error) {
	if owa.storage != nil {
		url := req.URL.RequestURI()
		cached, ok := owa.storage.GetCached(url)
		if ok {
			return []byte(cached), nil
		}
	}

	resp, err := owa.webClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if owa.storage != nil {
		url := req.URL.RequestURI()
		owa.storage.CacheFor(url, string(respData), CACHE_TIME)
	}
	return respData, nil
}

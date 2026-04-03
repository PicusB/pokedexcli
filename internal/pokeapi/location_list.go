package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c Client) ListLocations(url *string) (RespShallowLocations, error) {
	requestURL := "https://pokeapi.co/api/v2/location-area/"
	if url != nil {
		requestURL = *url
	}
	var resData []byte
	if cached, ok := c.pokeCache.Get(requestURL); ok {
		resData = cached
	} else {
		res, err := c.httpClient.Get(requestURL)
		if err != nil {
			return RespShallowLocations{}, err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return RespShallowLocations{}, err
		}
		if res.StatusCode > 299 {
			return RespShallowLocations{}, fmt.Errorf("Invalid Status Code: %d", res.StatusCode)
		}
		resData = body
		c.pokeCache.Add(requestURL, resData)
	}
	locations := RespShallowLocations{}
	umerr := json.Unmarshal(resData, &locations)
	if umerr != nil {
		return RespShallowLocations{}, umerr
	}
	return locations, nil
}

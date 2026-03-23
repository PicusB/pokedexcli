package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) ListLocations(url *string) (RespShallowLocations, error) {
	requestURL := "https://pokeapi.co/api/v2/location-area/"
	if url != nil {
		requestURL = *url
	}
	res, err := http.Get(requestURL)
	if err != nil {
		return RespShallowLocations{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return RespShallowLocations{}, fmt.Errorf("Invalid Status Code: %d", res.StatusCode)
	}
	if err != nil {
		return RespShallowLocations{}, err
	}
	locations := RespShallowLocations{}
	umerr := json.Unmarshal(body, &locations)
	if umerr != nil {
		return RespShallowLocations{}, umerr
	}
	return locations, nil
}

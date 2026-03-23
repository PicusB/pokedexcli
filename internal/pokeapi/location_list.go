package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) ListLocations(url *string) (RespShallowLocations, error) {
	res, err := http.Get(*url)
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

package main

import "fmt"

func commandMap(cfg *config) error {
	result, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)

	if err != nil {
		return err
	}
	cfg.nextURL = result.Next
	cfg.previousURL = result.Previous
	for _, location := range result.Results {
		fmt.Println((location.Name))
	}

	return nil
}

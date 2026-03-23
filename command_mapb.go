package main

import "fmt"

func commandMapb(cfg *config) error {
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	if *cfg.previousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	result, err := cfg.pokeapiClient.ListLocations(cfg.previousURL)

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

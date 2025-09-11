package cli

import (
	"errors"
	"fmt"

	"github.com/ManoloEsS/pokedex/internal/api"
)

func commandMapf(cfg *Config, parameter string) error {
	locationsResp, err := cfg.PokeapiClient.GetLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	printLocations(locationsResp)

	return nil
}

func commandMapb(cfg *Config, parameter string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you are already on the first page...")
	}

	locationsResp, err := cfg.PokeapiClient.GetLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	printLocations(locationsResp)

	return nil
}

func printLocations(locations api.RespShallowLocations) {
	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Println()
}

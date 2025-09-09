package cli

import (
	"errors"
	"fmt"

	"github.com/ManoloEsS/pokedex/internal/api"
)

func commandMapf(cfg *Config) error {
	locationsJson, err := cfg.PokeapiClient.GetLocationAreas(cfg.NextLocationsURL)
	if err != nil {
		return err
	}
	printLocations(&locationsJson)

	cfg.NextLocationsURL = locationsJson.Next
	cfg.PrevLocationsURL = locationsJson.Previous
	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you are already on the first page...")
	}
	locationsJson, err := cfg.PokeapiClient.GetLocationAreas(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}
	printLocations(&locationsJson)

	cfg.NextLocationsURL = locationsJson.Next
	cfg.PrevLocationsURL = locationsJson.Previous

	return nil
}

func printLocations(locations *api.RespShallowLocations) {
	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Println()
}

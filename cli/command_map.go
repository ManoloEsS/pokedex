package cli

import (
	"fmt"

	"github.com/ManoloEsS/pokedex/internal/api"
)

func commandMapf(cfg *Config) error {
	locationsJson, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	printLocations(&locationsJson)

	cfg.nextLocationsURL = locationsJson.Next
	cfg.prevLocationsURL = locationsJson.Previous
	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you are already on the first page...")
		return nil
	}
	cfg.nextLocationsURL = cfg.prevLocationsURL
	locationsJson, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	printLocations(&locationsJson)

	cfg.prevLocationsURL = locationsJson.Previous
	return nil
}

func printLocations(locations *api.RespShallowLocations) {
	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Println()
}

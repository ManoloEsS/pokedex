package cli

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ManoloEsS/pokedex/internal/api"
	"github.com/ManoloEsS/pokedex/internal/cache"
)

func commandMapf(cfg *Config, cache *cache.Cache) error {
	locationsData := &api.RespShallowLocations{}

	url := ""
	if cfg.NextLocationsURL != nil {
		url = *cfg.NextLocationsURL
	}

	if cached, found := cache.Get(url); found {
		err := json.Unmarshal(cached, locationsData)
		if err != nil {
		}
		return err
	} else {
		resp, err := cfg.PokeapiClient.GetLocationAreasAlt(cfg.NextLocationsURL, cache)
		if err != nil {
			return err
		}
		*locationsData = resp
	}

	printLocations(locationsData)

	cfg.NextLocationsURL = locationsData.Next
	cfg.PrevLocationsURL = locationsData.Previous

	return nil
}

func commandMapb(cfg *Config, cache *cache.Cache) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you are already on the first page...")
	}

	locationsData := &api.RespShallowLocations{}

	if cached, found := cache.Get(*cfg.PrevLocationsURL); found {
		err := json.Unmarshal(cached, locationsData)
		if err != nil {
			return err
		}
	} else {
		resp, err := cfg.PokeapiClient.GetLocationAreasAlt(cfg.PrevLocationsURL, cache)
		if err != nil {
			return err
		}
		*locationsData = resp
	}

	printLocations(locationsData)

	cfg.NextLocationsURL = locationsData.Next
	cfg.PrevLocationsURL = locationsData.Previous

	return nil
}

func printLocations(locations *api.RespShallowLocations) {
	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Println()
}

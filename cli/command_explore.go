package cli

import (
	"errors"
	"fmt"

	"github.com/ManoloEsS/pokedex/internal/api"
)

func commandExplore(cfg *Config, areaName string) error {
	if areaName == "" {
		return errors.New("Area name not provided. Usage 'explore <area name>'")
	}

	AreaData, err := cfg.PokeapiClient.GetAreaData(areaName)
	if err != nil {
		return err
	}

	printPokemon(AreaData)

	return nil
}

func printPokemon(data api.AreaData) {
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	fmt.Println()
}

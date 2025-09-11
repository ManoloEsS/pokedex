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

	areaData, err := cfg.PokeapiClient.GetAreaData(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	printPokemon(&areaData)

	return nil
}

func printPokemon(data *api.AreaData) {
	fmt.Println("Found Pokemon:")
	for _, e := range data.PokemonEncounters {
		fmt.Printf("- %s\n", e.Pokemon.Name)
	}
	fmt.Println()
}

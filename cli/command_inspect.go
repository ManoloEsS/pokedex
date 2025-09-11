package cli

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("Pokemon name not provided. Usage 'inspect <pokemon name>'")
	}

	if _, ok := cfg.Pokedex[pokemonName]; !ok {
		fmt.Println("you have not caught that pokemon")
	}
	pokemon := cfg.Pokedex[pokemonName]

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	return nil
}

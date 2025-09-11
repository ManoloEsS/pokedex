package cli

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("Pokemon name not provided. Usage 'inspect <pokemon name>'")
	}

	pokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}
	fmt.Println()

	return nil
}

package cli

import "fmt"

func commandPokedex(cfg *Config, parameter string) error {
	fmt.Println("Your Pokedex")
	for k := range cfg.Pokedex {
		fmt.Printf("  - %s\n", k)
	}
	fmt.Println()

	return nil
}

package cli

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/ManoloEsS/pokedex/internal/api"
)

func commandCatch(cfg *Config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("Pokemon name not provided. Usage 'catch <pokemon name>'")
	}

	pokemonData, err := cfg.PokeapiClient.GetPokemonData(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	if !catchAttemptSuccess(pokemonData) {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.addToPokedex(pokemonData)
	return nil
}

func catchAttemptSuccess(pokemon api.PokemonData) bool {
	maxLevel := 400

	if pokemon.BaseExperience > maxLevel {
		maxLevel = pokemon.BaseExperience + 1
	}

	failureChance := float64(pokemon.BaseExperience) / float64(maxLevel)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	roll := rng.Float64()

	return roll > failureChance
}

func (cfg *Config) addToPokedex(pokemon api.PokemonData) {
	if _, ok := cfg.Pokedex[pokemon.Name]; ok {
		return
	}

	cfg.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s has been added to the pokedex\n", pokemon.Name)
}

package main

import (
	"github.com/ManoloEsS/pokedex/cli"
	"github.com/ManoloEsS/pokedex/internal/api"
	"github.com/ManoloEsS/pokedex/internal/cache"
)

func main() {
	cfg := &cli.Config{
		PokeapiClient:    api.NewClient(0),
		NextLocationsURL: nil,
		PrevLocationsURL: nil,
	}

	cache := cache.NewCache(0)
	cli.StartRepl(cfg, cache)
}

package main

import (
	"github.com/ManoloEsS/pokedex/cli"
	"github.com/ManoloEsS/pokedex/internal/api"
)

func main() {
	cfg := &cli.Config{
		PokeapiClient:    api.NewClient(0),
		NextLocationsURL: nil,
		PrevLocationsURL: nil,
	}
	cli.StartRepl(cfg)
}

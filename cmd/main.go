package main

import (
	"time"

	"github.com/ManoloEsS/pokedex/cli"
	"github.com/ManoloEsS/pokedex/internal/api"
)

func main() {

	cfg := &cli.Config{
		PokeapiClient: api.NewClient(5*time.Second, 5*time.Minute),
	}
	cli.StartRepl(cfg)
}

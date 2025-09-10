package cli

import (
	"fmt"
	"os"

	"github.com/ManoloEsS/pokedex/internal/cache"
)

func commandExit(cfg *Config, cache *cache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

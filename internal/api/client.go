package api

import (
	"net/http"
	"time"

	"github.com/ManoloEsS/pokedex/internal/cache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      cache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      cache.NewCache(cacheInterval),
	}
}

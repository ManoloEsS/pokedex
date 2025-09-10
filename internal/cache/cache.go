package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    *map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval int) (*Cache, error) {

	return nil, nil
}


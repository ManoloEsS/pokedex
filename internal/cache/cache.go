package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheData map[string]cacheEntry
	mux       *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(intervalDuration time.Duration) Cache {
	cache := Cache{
		cacheData: make(map[string]cacheEntry),
		mux:       &sync.RWMutex{},
	}

	go cache.reapLoop(intervalDuration)

	return cache
}

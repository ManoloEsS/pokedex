package cache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheData map[string]CacheEntry
	Mu        sync.RWMutex
	Interval  time.Duration
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(intervalDuration time.Duration) *Cache {
	if intervalDuration <= 0 {
		intervalDuration = 5 * time.Second
	}
	cache := &Cache{CacheData: make(map[string]CacheEntry), Interval: intervalDuration}

	go cache.reapLoop()

	return cache
}

package cache

import (
	"time"
)

func (c *Cache) Add(key string, value []byte) {
	newEntry := CacheEntry{CreatedAt: time.Now(), Val: value}
	c.Mu.Lock()
	c.CacheData[key] = newEntry
	c.Mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.RLock()
	value, ok := c.CacheData[key]
	c.Mu.RUnlock()
	if !ok {
		return nil, false
	}
	return value.Val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()

	for range ticker.C {
		c.Mu.Lock()
		for key, entry := range c.CacheData {
			if time.Since(entry.CreatedAt) > c.Interval {
				delete(c.CacheData, key)
			}
		}
		c.Mu.Unlock()
	}
}

package cache

import (
	"time"
)

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	c.cacheData[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	value, ok := c.cacheData[key]
	c.mux.RUnlock()
	if !ok {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, entry := range c.cacheData {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.cacheData, key)
		}
	}
}

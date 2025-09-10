package cache

import "time"

func (d Cache) Add(key string, val []byte) error {
	return nil
}

func (d Cache) Get(key string) ([]byte, bool) {
	return nil, false
}

func (d Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(d.interval)
	for {
	}

}

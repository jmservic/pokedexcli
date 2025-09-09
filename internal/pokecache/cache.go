package pokecache

import (
	"time"
	"sync"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data: map[string]cacheEntry{},
		mu: sync.Mutex{},
	}
	cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			currTime := <-ticker.C
			c.mu.Lock()
			for key, entry := range c.data {
				elapsed := entry.createdAt.Sub(currTime)
				if elapsed >= interval {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}

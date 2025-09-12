package pokecache

import (
	"time"
	"sync"
	//"fmt"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: map[string]cacheEntry{},
		mu: &sync.Mutex{},
	}
	cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	//fmt.Printf("Pass by Pointer Mutex address: %p\n", (c.mu))
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	//fmt.Printf("Pass by Value Mutex address: %p\n", (c.mu))
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	tickerChan := time.Tick(interval)

	go func() {
		for {
			currTime := <-tickerChan
			//fmt.Println("\nReaping...\n")
			c.mu.Lock()
			for key, entry := range c.data {
				elapsed := currTime.Sub(entry.createdAt)
				//fmt.Println("%d time elapsed", elapsed) 
				if elapsed >= interval {
				//	fmt.Println("Reaping %s", key)
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}

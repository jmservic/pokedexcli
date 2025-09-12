package pokecache

import (
	"time"
	"sync"
)
type Cache struct {
	data map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

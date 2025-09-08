package pokecache

import (
	"time"
)
type Cache struct {
	interval time.Duration,
	data map[string]cacheEntry,
	//need mutex
}

type cacheEntry struct {
	createdAt time.Time,
	val []byte
}

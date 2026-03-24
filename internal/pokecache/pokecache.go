package pokecache

import "time"

type Cache struct {
	cacheEntry map[string]CacheEntry
	duration   time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

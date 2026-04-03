package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]CacheEntry
	duration     time.Duration
	lock         sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		cacheEntries: map[string]CacheEntry{},
		duration:     interval,
	}
	go newCache.reapLoop()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cacheEntries[key] = CacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	returnvalue, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return returnvalue.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration)
	defer ticker.Stop()
	for range ticker.C {
		c.lock.Lock()
		for index, cachedEntry := range c.cacheEntries {
			age := time.Since(cachedEntry.createdAt)
			if age > c.duration {
				delete(c.cacheEntries, index)
			}
		}
		c.lock.Unlock()
	}
}

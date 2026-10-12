package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	cacheMu  sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheMap: map[string]cacheEntry{},
		cacheMu:  sync.Mutex{},
	}
	cache.reaploop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	cacheEntry := cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()
	c.cacheMap[key] = cacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()

	cacheEntry, ok := c.cacheMap[key]

	if !ok {
		return []byte{}, false
	}

	value := cacheEntry.value
	return value, true
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop() // Stops the reap if eventually the function return

	for range ticker.C {
		c.cacheMu.Lock()
		for key, val := range c.cacheMap {
			if time.Since(val.createdAt) > interval {
				delete(c.cacheMap, key)
			}
		}
		c.cacheMu.Unlock()
	}
}

package pokeCache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entry map[string]CacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	//create and initialize new cache
	var newCache Cache
	newCache.entry = make(map[string]CacheEntry)

	//run cache cleanup loop
	go newCache.reapLoop(interval)
	return &newCache
}

// add a record to cache
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	c.entry[key] = newEntry
}

// get a record from cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.entry[key].val) == 0 {
		return nil, false
	}
	return c.entry[key].val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	for range ticker.C {
		for k, cEntry := range c.entry {
			cacheCutover := time.Now().Add(-interval)
			if cacheCutover.After(cEntry.createdAt) {
				c.mu.Lock()
				delete(c.entry, k)
				c.mu.Unlock()
			}
		}
	}
}

func (c *Cache) Clear() {
	c.mu.Lock()
	c.entry = make(map[string]CacheEntry)
	c.mu.Unlock()
}

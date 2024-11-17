package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Value     []byte
	TTL       int
}

func (ce *CacheEntry) HasExpired() bool {
	return time.Now().After(ce.CreatedAt.Add(time.Second * time.Duration(ce.TTL)))
}

type Cache struct {
	Entries map[string]CacheEntry
	mu      sync.Mutex
}

func (c *Cache) Add(key string, val []byte, ttl int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := CacheEntry{
		CreatedAt: time.Now(),
		Value:     val,
		TTL:       ttl,
	}

	c.Entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, found := c.Entries[key]
	if !found {
		return nil, false
	}

	if entry.HasExpired() {
		delete(c.Entries, key)
		return nil, false
	}

	return entry.Value, true
}

func (c *Cache) Delete(key string) {
	delete(c.Entries, key)
}

func (c *Cache) Cleanup() {
	for range time.Tick(time.Second * 10) {
		c.mu.Lock()

		for index, entry := range c.Entries {
			if entry.HasExpired() {
				delete(c.Entries, index)
			}
		}

		c.mu.Unlock()
	}
}

var cacheInstance *Cache

func CacheInstance() *Cache {
	if cacheInstance == nil {
		cacheInstance = &Cache{
			Entries: make(map[string]CacheEntry),
		}

		go cacheInstance.Cleanup()
	}

	return cacheInstance
}

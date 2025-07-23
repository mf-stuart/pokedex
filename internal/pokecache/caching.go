package pokecache

import (
	"sync"
	"time"
)

const DEFAULT_INTERVAL = time.Duration(10 * time.Second)

var LocationCache = NewCache(DEFAULT_INTERVAL)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cacheable interface {
	Add(key string, value []byte)
	Get(key string) ([]byte, bool)
}
type Cache struct {
	mu      sync.RWMutex
	entries map[string]*cacheEntry
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, exists := c.entries[key]
	if !exists {
		c.entries[key] = &cacheEntry{createdAt: time.Now(), val: value}
	}
	return
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			if entry.createdAt.Add(interval).Before(now) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{entries: make(map[string]*cacheEntry)}
	go newCache.reapLoop(interval)
	return &newCache
}

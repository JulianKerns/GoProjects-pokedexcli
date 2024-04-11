package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu   *sync.Mutex
	cmap map[string]CacheEntry
}
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	tInitial := time.Now()
	c.cmap[key] = CacheEntry{createdAt: tInitial, val: value}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.cmap[key]
	if !ok {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) ReapLoop(intervalTime time.Duration) {
	ticker := time.NewTicker(intervalTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()

			for key := range c.cmap {
				t0 := time.Now()
				if t0.After(c.cmap[key].createdAt.Add(intervalTime)) {
					delete(c.cmap, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cacheMap := make(map[string]CacheEntry)
	cacheMutex := &sync.Mutex{}
	InitialCache := &Cache{mu: cacheMutex, cmap: cacheMap}
	go InitialCache.ReapLoop(interval)
	return InitialCache
}

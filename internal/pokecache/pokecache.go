package pokecache

import (
	"fmt"
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
	tInitial := time.Now()
	c.cmap[key] = CacheEntry{createdAt: tInitial, val: value}
}

func (c *Cache) Get(key string) ([]byte, bool) {
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

func NewCache(interval time.Duration) {
	t0 := time.Now()
	//cacheMap := make(map[string]CacheEntry)
	//cacheMutex := &sync.Mutex{}
	//InitialCache := Cache{mu: cacheMutex, cmap: cacheMap}
	t1 := time.Now()
	fmt.Printf("Initialization took %v miliseconds.\n", t1.Sub(t0))
}

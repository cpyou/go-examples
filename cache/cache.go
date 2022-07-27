package cache

import (
	"sync"
	"time"
)

const (
	// NeverClean 永远不会自动清除所有缓存
	NeverClean time.Duration = -1
)

type Cache struct {
	items    map[string]interface{}
	capacity uint32
	mu       sync.RWMutex
}

func NewCache(capacity uint32, cleanupInterval time.Duration) *Cache {
	c := &Cache{
		items:    make(map[string]interface{}),
		capacity: capacity,
	}

	if cleanupInterval > 0 {
		go func() {
			ticker := time.NewTicker(cleanupInterval)
			defer ticker.Stop()

			for range ticker.C {
				c.reset()
			}
		}()
	}

	return c
}

func (c *Cache) Set(key string, val interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.items) == int(c.capacity) {
		c.reset()
	}

	c.items[key] = val
}

func (c *Cache) Get(key string) (val interface{}, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	return item, found
}

func (c *Cache) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]interface{})
}

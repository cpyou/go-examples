package cache

import (
	"sync"
	"time"
)

type TimesCache struct {
	mu   sync.RWMutex
	Data map[string]uint32
}

func NewTimesCounter(cleanupInterval time.Duration) *TimesCache {
	c := &TimesCache{
		mu:   sync.RWMutex{},
		Data: make(map[string]uint32),
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

func (c *TimesCache) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Data = make(map[string]uint32)
}

func (c *TimesCache) get(key string) uint32 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if val, ok := c.Data[key]; ok {
		return val
	}
	return 0
}

func (c *TimesCache) Incr(key string, val uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.Data[key]
	if ok {
		c.Data[key] += val
	} else {
		c.Data[key] = val
	}
	return
}

func (c *TimesCache) decr(key string, val uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.Data[key]
	if ok {
		c.Data[key] -= val
	} else {
		c.Data[key] = 0
	}
	return
}

func (c *TimesCache) remove(field string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.Data, field)
	return
}

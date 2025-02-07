package cache

import (
	"sync"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type InMemoryCache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		store: make(map[string]interface{}),
	}
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.store[key]
	return value, exists
}

func (c *InMemoryCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}
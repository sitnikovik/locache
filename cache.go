package locache

import (
	"sync"
)

// Cache is a simple in-memory key-value storage optimized for concurrent use.
// The cache is safe for concurrent use by multiple goroutines.
// The cache can be used to store any type of value.
type Cache struct {
	// m - map of key value pairs
	m map[string]any
	// mu - mutex to lock the map
	mu sync.RWMutex
}

// NewCache creates a new cache with a default time to live of 0.
func NewCache() *Cache {
	return &Cache{
		m: make(map[string]any),
	}
}

// Get returns the value for the given key.
// If the key is not found, the second return value is false.
// If the key is found, the second return value is true.
func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.m[key]
	return value, ok
}

// Add adds the key value pair to the cache if the key does not exist.
// If the key already exists, it does nothing.
func (c *Cache) Add(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[key]; !ok {
		c.m[key] = value
	}
}

// Set sets the value for the given key in the cache even if the key already exists.
func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value
}

// Has returns true if the key exists in the cache.
func (c *Cache) Has(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.m[key]
	return ok
}

// Delete removes the key from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.m, key)
}

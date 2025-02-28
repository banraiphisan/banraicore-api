package inmem

import (
	"errors"
	"sync"
	"time"
)

type InMemoryCache struct {
	mu    sync.RWMutex
	store map[string]Item
}

type Item struct {
	Value      []byte
	Expiration int64
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		store: make(map[string]Item),
	}
}

func (c *InMemoryCache) Set(key string, value []byte, duration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	expiration := time.Now().Add(duration).UnixNano()
	c.store[key] = Item{
		Value:      value,
		Expiration: expiration,
	}

	return nil
}

func (c *InMemoryCache) Get(key string) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.store[key]
	if !found || time.Now().UnixNano() > item.Expiration {
		return nil, errors.New("cache miss")
	}

	return item.Value, nil
}

func (c *InMemoryCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)

	return nil
}

func (c *InMemoryCache) Close() error {
	return nil
}

func (c *InMemoryCache) Reset() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store = make(map[string]Item)

	return nil
}

func (c *InMemoryCache) Ping() error {
	return nil
}

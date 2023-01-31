package mutex

import (
	"study/cmd/cache/storage"
	"sync"
)

type SimpleCache struct {
	storage map[string]string
	mu      sync.Mutex
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{storage: make(map[string]string)}
}

func (cache *SimpleCache) Get(key string) (string, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	value, ok := cache.storage[key]
	if !ok {
		return "", storage.ErrNotFound
	}
	return value, nil
}

func (cache *SimpleCache) Set(key, value string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.storage[key] = value
	return nil
}

func (cache *SimpleCache) Delete(key string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	delete(cache.storage, key)
	return nil
}

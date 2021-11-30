package domain

import (
	"jarjarbinks/pkg/domain/repository"
	"sync"
	"time"
)

type cacheStore struct {
	mu     sync.RWMutex
	unsafe repository.CacheStore
}

func (c *cacheStore) Load(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.unsafe.Load(key)
}

func (c *cacheStore) Peek(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.unsafe.Peek(key)
}

func (c *cacheStore) Update(key interface{}, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsafe.Update(key, value)
}

func (c *cacheStore) Store(key interface{}, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsafe.Store(key, value)
}

func (c *cacheStore) StoreWithTTL(key interface{}, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsafe.StoreWithTTL(key, value, ttl)
}

func (c *cacheStore) Delete(key interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsafe.Delete(key)
}

func (c *cacheStore) Keys() []interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.unsafe.Keys()
}

func (c *cacheStore) Contains(key interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.unsafe.Contains(key)
}

func (c *cacheStore) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsafe.Purge()
}

func (c *cacheStore) Resize(s int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.unsafe.Resize(s)
}

func (c *cacheStore) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.unsafe.Len()
}

func (c *cacheStore) Cap() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.unsafe.Cap()
}

func (c *cacheStore) TTL() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.unsafe.TTL()
}

func (c *cacheStore) SetTTL(ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsafe.SetTTL(ttl)
}

func (c *cacheStore) RegisterOnEvicted(f func(key, value interface{})) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsafe.RegisterOnEvicted(f)
}

func (c *cacheStore) RegisterOnExpired(f func(key, value interface{})) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.unsafe.RegisterOnExpired(f)
}

func (c *cacheStore) Expiry(key interface{}) (time.Time, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.unsafe.Expiry(key)
}

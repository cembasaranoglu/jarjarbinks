package domain

import (
	"time"
)

type Cache struct {
	coll    Collection
	entries map[interface{}]*Entry
	onEvicted func(key, value interface{})
	onExpired func(key, value interface{})
	ttl       time.Duration
	capacity  int
}


func (c *Cache) Load(key interface{}) (interface{}, bool) {
	return c.get(key, false)
}

func (c *Cache) Peek(key interface{}) (interface{}, bool) {
	return c.get(key, true)
}

func (c *Cache) get(key interface{}, peek bool) (v interface{}, found bool) {
	e, ok := c.entries[key]
	if !ok {
		return
	}

	if !e.Exp.IsZero() && time.Now().UTC().After(e.Exp) {
		c.evict(e)
		return
	}

	if !peek {
		c.coll.Move(e)
	}

	return e.Value, ok
}

func (c *Cache) Expiry(key interface{}) (t time.Time, ok bool) {
	ok = c.Contains(key)
	if ok {
		t = c.entries[key].Exp
	}
	return t, ok
}

func (c *Cache) Store(key, value interface{}) {
	c.StoreWithTTL(key, value, c.ttl)
}
func (c *Cache) StoreWithTTL(key, value interface{}, ttl time.Duration) {
	if e, ok := c.entries[key]; ok {
		c.removeEntry(e)
	}

	e := &Entry{Key: key, Value: value}

	if ttl > 0 {
		if c.onExpired != nil {
			e.startTimer(ttl, c.onExpired)
		}
		e.Exp = time.Now().UTC().Add(ttl)
	}

	c.entries[key] = e
	if c.capacity != 0 && c.Len() >= c.capacity {
		c.Discard()
	}
	c.coll.Add(e)
}

func (c *Cache) Update(key, value interface{}) {
	if c.Contains(key) {
		c.entries[key].Value = value
	}
}

func (c *Cache) Purge() {
	defer c.coll.Init()

	if c.onEvicted == nil {
		c.entries = make(map[interface{}]*Entry)
		return
	}

	for _, e := range c.entries {
		c.evict(e)
	}
}

func (c *Cache) Resize(size int) int {
	c.capacity = size
	diff := c.Len() - size

	if diff < 0 {
		diff = 0
	}

	for i := 0; i < diff; i++ {
		c.Discard()
	}

	return diff
}

func (c *Cache) DelSilently(key interface{}) {
	if e, ok := c.entries[key]; ok {
		c.removeEntry(e)
	}
}

func (c *Cache) Delete(key interface{}) {
	if e, ok := c.entries[key]; ok {
		c.evict(e)
	}
}

// Contains Checks if a key exists in cache.
func (c *Cache) Contains(key interface{}) (ok bool) {
	_, ok = c.Peek(key)
	return
}

// Keys return cache records keys.
func (c *Cache) Keys() (keys []interface{}) {
	for k := range c.entries {
		keys = append(keys, k)
	}
	return
}

func (c *Cache) Len() int {
	return c.coll.Len()
}

func (c *Cache) Discard() (key, value interface{}) {
	if e := c.coll.Discard(); e != nil {
		c.evict(e)
		return e.Key, e.Value
	}

	return
}

func (c *Cache) removeEntry(e *Entry) {
	c.coll.Remove(e)
	e.stopTimer()
	delete(c.entries, e.Key)
}

func (c *Cache) evict(e *Entry) {
	c.removeEntry(e)
	if c.onEvicted != nil {
		go c.onEvicted(e.Key, e.Value)
	}
}

func (c *Cache) TTL() time.Duration {
	return c.ttl
}

func (c *Cache) SetTTL(ttl time.Duration) {
	c.ttl = ttl
}

func (c *Cache) Cap() int {
	return c.capacity
}
func (c *Cache) RegisterOnEvicted(f func(key, value interface{})) {
	c.onEvicted = f
}

func (c *Cache) RegisterOnExpired(f func(key, value interface{})) {
	c.onExpired = f
}

func New(c Collection, cap int) *Cache {
	return &Cache{
		coll:     c,
		capacity: cap,
		entries:  make(map[interface{}]*Entry),
	}
}
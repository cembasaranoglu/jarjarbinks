package repository

import "time"

type CacheStore interface {
	Load(key interface{}) (interface{}, bool)
	Peek(key interface{}) (interface{}, bool)
	Update(key interface{}, value interface{})
	Store(key interface{}, value interface{})
	StoreWithTTL(key interface{}, value interface{}, ttl time.Duration)
	Delete(key interface{})
	Expiry(key interface{}) (time.Time, bool)
	Keys() []interface{}
	Contains(key interface{}) bool
	Purge()
	Resize(int) int
	Len() int
	Cap() int
	TTL() time.Duration
	SetTTL(time.Duration)
	RegisterOnEvicted(f func(key, value interface{}))
	RegisterOnExpired(f func(key, value interface{}))
}

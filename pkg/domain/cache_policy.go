package domain

import (
	"errors"
	"jarjarbinks/pkg/domain/repository"
	"strconv"
	"sync"
)

type CachePolicy uint
const (
	IDLE CachePolicy = iota + 1
	FIFO
	LIFO
	LRU
	LFU
	MRU
	ARC
	max
)
func (c CachePolicy) String() string {
	switch c {
	case IDLE:
		return "IDLE"
	case FIFO:
		return "FIFO"
	case LIFO:
		return "LIFO"
	case LRU:
		return "LRU"
	case LFU:
		return "LFU"
	case MRU:
		return "MRU"
	case ARC:
		return "ARC"
	default:
		return "unknown cache replacement policy value " + strconv.Itoa(int(c))
	}

}

var policies = make([]func(cap int) repository.CacheStore, max)
func (c CachePolicy) Register(function func(cap int)  repository.CacheStore) error{
	if c <= 0 && c >= max {
		return errors.New("cache policy does not support")
	}

	policies[c] = function
	return nil
}
func (c CachePolicy) Available() bool {
	return c > 0 && c < max && policies[c] != nil
}

func (c CachePolicy) New(cap int) repository.CacheStore {
	cache := new(cacheStore)
	cache.mu = sync.RWMutex{}
	cache.unsafe = c.NewUnsafe(cap)
	return cache
}

func (c CachePolicy) NewUnsafe(cap int) repository.CacheStore {
	if !c.Available() {
		panic("policy is not available")
	}

	return policies[c](cap)
}
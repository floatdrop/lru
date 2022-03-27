// Package lru implements cache with least recent used eviction policy.
package lru

import (
	"sync"

	list "github.com/bahlo/generic-list-go"
)

// LRU implements Cache interface with least recent used eviction policy.
type LRU[K comparable, V any] struct {
	m     sync.Mutex
	ll    *list.List[*entry[K, V]]
	cache map[K]*list.Element[*entry[K, V]]
	size  int
}

type entry[K comparable, V any] struct {
	key   K
	value *V
}

// Get returns pointer to value for key, if value was in cache (nil returned otherwise).
func (L *LRU[K, V]) Get(key K) *V {
	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.cache[key]; ok {
		L.ll.MoveToFront(e)
		return e.Value.value
	}

	return nil
}

// Set inserts key value pair and returns evicted value, if cache was full.
// If cache size is less than 1 – method will always return reference to value (as if it was immediately evicted).
func (L *LRU[K, V]) Set(key K, value V) *V {
	if L.size < 1 {
		return &value
	}

	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.cache[key]; ok {
		previousValue := e.Value.value
		L.ll.MoveToFront(e)
		e.Value.value = &value
		return previousValue
	}

	e := L.ll.Back()
	i := e.Value
	evictedValue := i.value
	delete(L.cache, i.key)

	i.key = key
	i.value = &value
	L.cache[key] = e
	L.ll.MoveToFront(e)
	return evictedValue
}

// Len returns number of cached items.
func (L *LRU[K, V]) Len() int {
	L.m.Lock()
	defer L.m.Unlock()

	return len(L.cache)
}

// Remove method removes entry associated with key and returns pointer to removed value (or nil if entry was not in cache).
func (L *LRU[K, V]) Remove(key K) *V {
	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.cache[key]; ok {
		value := e.Value.value
		L.ll.MoveToBack(e)
		e.Value.value = nil
		delete(L.cache, key)
		return value
	}

	return nil
}

// Peek returns value for key (if key was in cache), but does not modify its recency.
func (L *LRU[K, V]) Peek(key K) *V {
	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.cache[key]; ok {
		return e.Value.value
	}

	return nil
}

// New creates LRU cache with size capacity. Cache will preallocate size count of internal structures to avoid allocation in process.
func New[K comparable, V any](size int) *LRU[K, V] {
	c := &LRU[K, V]{
		ll:    list.New[*entry[K, V]](),
		cache: make(map[K]*list.Element[*entry[K, V]], size),
		size:  size,
	}

	for i := 0; i < size; i++ {
		c.ll.PushBack(&entry[K, V]{})
	}

	return c
}

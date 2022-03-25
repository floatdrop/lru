// Package lru implements cache with least recent used eviction policy.
package lru

import (
	"sync"

	list "github.com/bahlo/generic-list-go"
)

// LRU implements Cache interface with least recent used eviction policy.
type LRU[K comparable, V any] struct {
	m        sync.Mutex
	list     *list.List[*entry[K, V]]
	elements map[K]*list.Element[*entry[K, V]]
	size     int
}

type entry[K comparable, V any] struct {
	key   K
	value *V
}

// Get returns pointer to value for key, if value was in cache (nil returned otherwise).
func (L *LRU[K, V]) Get(key K) *V {
	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.elements[key]; ok {
		L.list.MoveToFront(e)
		return e.Value.value
	}

	return nil
}

// Set inserts key value pair and returns evicted value, if cache was full.
func (L *LRU[K, V]) Set(key K, value V) *V {
	if L.size == 0 {
		return nil
	}

	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.elements[key]; ok {
		previousValue := e.Value.value
		L.list.MoveToFront(e)
		e.Value.value = &value
		return previousValue
	}

	e := L.list.Back()
	i := e.Value
	evictedValue := i.value
	delete(L.elements, i.key)

	i.key = key
	i.value = &value
	L.elements[key] = e
	L.list.MoveToFront(e)
	return evictedValue
}

// New creates LRU cache with size capacity. Cache will preallocate size count of internal structures to avoid allocation in process.
func New[K comparable, V any](size int) *LRU[K, V] {
	c := &LRU[K, V]{
		elements: make(map[K]*list.Element[*entry[K, V]], size),
		list:     list.New[*entry[K, V]](),
		size:     size,
	}

	for i := 0; i < size; i++ {
		c.list.PushFront(&entry[K, V]{})
	}

	return c
}

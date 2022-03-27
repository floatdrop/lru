package lru

// Cache defines minimal interface for all cache implementations.
type Cache[K comparable, V any] interface {
	Get(key K) *V

	Set(key K, value V) *Evicted[K, V]

	Len() int

	Remove(key K) *V

	Peek(key K) *V
}
